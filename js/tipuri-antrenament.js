let currentEditID = 0;
let categoriiList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    loadCategorii();

    document.getElementById('form-add-tip').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-tipuri').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const tipID = event.target.getAttribute('data-id');
            handleDelete(tipID);
        }
        if (event.target.classList.contains('btn-edit')) {
            const tipID = event.target.getAttribute('data-id');
            handleEditClick(tipID);
        }
    });
});

function loadCategorii() {
    fetch('/api/tipuri-antrenament')
        .then(response => response.json())
        .then(data => {
            categoriiList = data;
            renderCategorii();
        })
        .catch(error => console.error('Eroare la preluarea categoriilor:', error));
}

function renderCategorii() {
    const tbody = document.getElementById('lista-tipuri');
    tbody.innerHTML = '';
    if (!categoriiList) categoriiList = [];
    categoriiList.forEach(tip => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${tip.id}</td>
            <td>${tip.nume}</td>
            <td>${tip.descriere || '-'}</td>
            <td>
                <button class="btn-edit" data-id="${tip.id}">Editează</button>
                <button class="btn-delete" data-id="${tip.id}">Șterge</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function sortCategorii(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    categoriiList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        // Handle nulls safely
        if (valA === null || valA === undefined) valA = '';
        if (valB === null || valB === undefined) valB = '';

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderCategorii();
}

function handleFormSubmit(event) {
    event.preventDefault();

    const tipData = {
        nume: document.getElementById('nume-tip').value,
        descriere: document.getElementById('descriere-tip').value
    };

    let url = '/api/tipuri-antrenament/add';

    if (currentEditID !== 0) {
        tipData.id = currentEditID;
        url = '/api/tipuri-antrenament/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(tipData)
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la salvarea categoriei'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            resetFormular();
            loadCategorii();
        })
        .catch(error => console.error('Eroare formular:', error));
}

function handleEditClick(id) {
    fetch(`/api/tip-antrenament?id=${id}`)
        .then(response => response.json())
        .then(tip => {
            document.getElementById('nume-tip').value = tip.nume;
            document.getElementById('descriere-tip').value = tip.descriere || '';

            currentEditID = tip.id;
            document.querySelector('#form-add-tip button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor categoriei:', error));
}

function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi categoria cu ID-ul ${id}?`)) {
        return;
    }

    fetch('/api/tipuri-antrenament/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la ștergerea categoriei'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            loadCategorii();
        })
        .catch(error => console.error('Eroare la ștergere:', error));
}

function resetFormular() {
    document.getElementById('form-add-tip').reset();
    currentEditID = 0;
    document.querySelector('#form-add-tip button[type="submit"]').textContent = 'Adaugă Categorie';
}