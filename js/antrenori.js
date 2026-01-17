let currentEditID = 0;
let antrenoriList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    loadAntrenori();

    document.getElementById('form-add-antrenor').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-antrenori').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const antrenorID = event.target.getAttribute('data-id');
            handleDelete(antrenorID);
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
            antrenoriList = data;
            renderAntrenori();
        })
        .catch(error => alert('Eroare la preluarea antrenorilor: ' + error.message));
}

function renderAntrenori() {
    const tbody = document.getElementById('lista-antrenori');
    tbody.innerHTML = '';
    if (!antrenoriList) antrenoriList = [];
    antrenoriList.forEach(antrenor => {
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
}

function sortAntrenori(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    antrenoriList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderAntrenori();
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
            alert(data.mesaj);
            resetFormular();
            loadAntrenori();
        })
        .catch(error => alert('Eroare formular: ' + error.message));
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
        .catch(error => alert('Eroare: ' + error.message));
}

function handleDelete(id) {
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
            alert(data.mesaj);
            loadAntrenori();
        })
        .catch(error => alert('Eroare: ' + error.message));
}

function resetFormular() {
    document.getElementById('form-add-antrenor').reset();
    currentEditID = 0;
    document.querySelector('#form-add-antrenor button[type="submit"]').textContent = 'Adaugă Antrenor';
}