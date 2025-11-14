let currentEditID = 0;

window.addEventListener('DOMContentLoaded', () => {
    loadMembri();
    loadAbonamente();
    document.getElementById('form-add-membru').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-membri').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const membruID = event.target.getAttribute('data-id');
            handleDeleteMembru(membruID);
        }

        if (event.target.classList.contains('btn-edit')) {
            const membruID = event.target.getAttribute('data-id');
            handleEditClick(membruID);
        }
    });
});

function loadMembri() {
    fetch('/api/membri')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('lista-membri');
            tbody.innerHTML = '';
            data.forEach(membru => {
                const tr = document.createElement('tr');

                tr.innerHTML = `
                    <td>${membru.id}</td>
                    <td>${membru.nume}</td>
                    <td>${membru.prenume}</td>
                    <td>${membru.email}</td>
                    <td>
                        <button class="btn-edit" data-id="${membru.id}">Editează</button>
                        <button class="btn-delete" data-id="${membru.id}">Șterge</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea membrilor:', error));
}

function loadAbonamente() {
    fetch('/api/abonamente')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-abonament');
            select.innerHTML = '<option value="">Alege un abonament...</option>';
            data.forEach(abonament => {
                const option = document.createElement('option');
                option.value = abonament.id;
                option.textContent = `${abonament.tip} - ${abonament.pret} RON`;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare la preluarea abonamentelor:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();

    const membruData = {
        nume: document.getElementById('nume').value,
        prenume: document.getElementById('prenume').value,
        email: document.getElementById('email').value,
        abonamentID: parseInt(document.getElementById('select-abonament').value, 10)
    };

    if (currentEditID === 0) {
        fetch('/api/membri/add', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(membruData)
        })
            .then(response => {
                if (!response.ok) { throw new Error('Eroare la adăugarea membrului'); }
                return response.json();
            })
            .then(data => {
                console.log(data.mesaj);
                resetFormular();
                loadMembri();
            })
            .catch(error => console.error('Eroare formular adăugare:', error));
    } else {
        membruData.id = currentEditID;

        fetch('/api/membri/update', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(membruData)
        })
            .then(response => {
                if (!response.ok) { throw new Error('Eroare la actualizarea membrului'); }
                return response.json();
            })
            .then(data => {
                console.log(data.mesaj);
                resetFormular();
                loadMembri();
            })
            .catch(error => console.error('Eroare formular actualizare:', error));
    }
}

function handleEditClick(id) {
    fetch(`/api/membru?id=${id}`)
        .then(response => response.json())
        .then(membru => {
            document.getElementById('nume').value = membru.nume;
            document.getElementById('prenume').value = membru.prenume;
            document.getElementById('email').value = membru.email;
            document.getElementById('select-abonament').value = membru.abonamentID;

            currentEditID = membru.id;

            document.querySelector('#form-add-membru button[type="submit"]').textContent = 'Salvează Modificările';

            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor membrului:', error));
}

function resetFormular() {
    document.getElementById('form-add-membru').reset();
    currentEditID = 0;
    document.querySelector('#form-add-membru button[type="submit"]').textContent = 'Adaugă Membru';
}

function handleDeleteMembru(id) {
    if (!confirm(`Ești sigur că vrei să ștergi membrul cu ID-ul ${id}? Acțiunea este ireversibilă!`)) {
        return;
    }

    fetch('/api/membri/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la ștergerea membrului'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            loadMembri();
        })
        .catch(error => console.error('Eroare la ștergere:', error));
}