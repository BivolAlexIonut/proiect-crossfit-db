let currentEditID = 0;

window.addEventListener('DOMContentLoaded', () => {
    loadEchipamente();
    document.getElementById('form-add-echipament').addEventListener('submit', handleFormSubmit);
    document.getElementById('lista-echipamente').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            handleDelete(event.target.getAttribute('data-id'));
        }
        if (event.target.classList.contains('btn-edit')) {
            handleEditClick(event.target.getAttribute('data-id'));
        }
    });
});

function loadEchipamente() {
    fetch('/api/echipamente')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('lista-echipamente');
            tbody.innerHTML = '';
            data.forEach(echipament => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${echipament.id}</td>
                    <td>${echipament.nume}</td>
                    <td>${echipament.cantitate} buc.</td>
                    <td>
                        <button class="btn-edit" data-id="${echipament.id}">Editează</button>
                        <button class="btn-delete" data-id="${echipament.id}">Șterge</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea echipamentelor:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();
    const echipamentData = {
        nume: document.getElementById('nume-echipament').value,
        cantitate: parseInt(document.getElementById('cantitate-echipament').value, 10)
    };

    let url = '/api/echipamente/add';
    if (currentEditID !== 0) {
        echipamentData.id = currentEditID;
        url = '/api/echipamente/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(echipamentData)
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la salvarea echipamentului'); } return response.json(); })
        .then(data => { console.log(data.mesaj); resetFormular(); loadEchipamente(); })
        .catch(error => console.error('Eroare formular:', error));
}

function handleEditClick(id) {
    fetch(`/api/echipament?id=${id}`)
        .then(response => response.json())
        .then(echipament => {
            document.getElementById('nume-echipament').value = echipament.nume;
            document.getElementById('cantitate-echipament').value = echipament.cantitate;
            currentEditID = echipament.id;
            document.querySelector('#form-add-echipament button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor echipamentului:', error));
}

function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi echipamentul cu ID-ul ${id}?`)) { return; }
    fetch('/api/echipamente/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la ștergerea echipamentului'); } return response.json(); })
        .then(data => { console.log(data.mesaj); loadEchipamente(); })
        .catch(error => console.error('Eroare la ștergere:', error));
}

function resetFormular() {
    document.getElementById('form-add-echipament').reset();
    currentEditID = 0;
    document.querySelector('#form-add-echipament button[type="submit"]').textContent = 'Adaugă Echipament';
}