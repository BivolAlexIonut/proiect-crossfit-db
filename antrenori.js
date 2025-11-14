let currentEditID = 0;

window.addEventListener('DOMContentLoaded', () => {
    loadAntrenori();

    document.getElementById('form-add-antrenor').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-antrenori').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const antrenorID = event.target.getAttribute('data-id');
            handleDeleteAntrenor(antrenorID);
        }
        if (event.target.classList.contains('btn-edit')) {
            const antrenorID = event.target.getAttribute('data-id');
            handleEditClick(antrenorID);
        }
    });
});

function loadAntrenori() {
    fetch('/api/antrenori')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('lista-antrenori');
            tbody.innerHTML = '';
            data.forEach(antrenor => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${antrenor.id}</td>
                    <td>${antrenor.nume}</td>
                    <td>${antrenor.prenume}</td>
                    <td>${antrenor.specializare}</td>
                    <td>
                        <button class="btn-edit" data-id="${antrenor.id}">Editează</button>
                        <button class="btn-delete" data-id="${antrenor.id}">Șterge</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea antrenorilor:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();

    const antrenorData = {
        nume: document.getElementById('nume-antrenor').value,
        prenume: document.getElementById('prenume-antrenor').value,
        specializare: document.getElementById('specializare-antrenor').value
    };

    let url = '/api/antrenori/add';

    if (currentEditID !== 0) {
        antrenorData.id = currentEditID;
        url = '/api/antrenori/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(antrenorData)
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la salvarea antrenorului'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            resetFormular();
            loadAntrenori();
        })
        .catch(error => console.error('Eroare formular:', error));
}

function handleEditClick(id) {
    fetch(`/api/antrenor?id=${id}`)
        .then(response => response.json())
        .then(antrenor => {
            document.getElementById('nume-antrenor').value = antrenor.nume;
            document.getElementById('prenume-antrenor').value = antrenor.prenume;
            document.getElementById('specializare-antrenor').value = antrenor.specializare;

            currentEditID = antrenor.id;
            document.querySelector('#form-add-antrenor button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor antrenorului:', error));
}

function handleDeleteAntrenor(id) {
    if (!confirm(`Ești sigur că vrei să ștergi antrenorul cu ID-ul ${id}?`)) {
        return;
    }

    fetch('/api/antrenori/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la ștergerea antrenorului'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            loadAntrenori();
        })
        .catch(error => console.error('Eroare la ștergere:', error));
}

function resetFormular() {
    document.getElementById('form-add-antrenor').reset();
    currentEditID = 0;
    document.querySelector('#form-add-antrenor button[type="submit"]').textContent = 'Adaugă Antrenor';
}