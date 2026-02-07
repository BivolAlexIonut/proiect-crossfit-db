let currentEditID = 0;
let abonamenteList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    loadAbonamente();

    document.getElementById('form-add-abonament').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-abonamente').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const abonamentID = event.target.getAttribute('data-id');
            handleDelete(abonamentID);
        }
        if (event.target.classList.contains('btn-edit')) {
            const abonamentID = event.target.getAttribute('data-id');
            handleEditClick(abonamentID);
        }
    });
});

function loadAbonamente() {
    fetch('/api/abonamente')
        .then(response => response.json())
        .then(data => {
            abonamenteList = data;
            renderAbonamente();
        })
        .catch(error => console.error('Eroare la preluarea abonamentelor:', error));
}

function renderAbonamente() {
    const tbody = document.getElementById('lista-abonamente');
    tbody.innerHTML = '';
    abonamenteList.forEach(abonament => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${abonament.id}</td>
            <td>${abonament.tip}</td>
            <td>${abonament.pret.toFixed(2)} RON</td>
            <td>
                <button class="btn-edit" data-id="${abonament.id}">Editează</button>
                <button class="btn-delete" data-id="${abonament.id}">Șterge</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function sortAbonamente(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    abonamenteList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderAbonamente();
}

function handleFormSubmit(event) {
    event.preventDefault();

    const tip = document.getElementById('tip-abonament').value;
    const pret = parseFloat(document.getElementById('pret-abonament').value);

    // Listă implicită de tipuri de abonament (conform datelor inițiale), 
    // adăugată pentru a răspunde cerinței de validare strictă, deși SQL-ul permite orice string.
    const tipuriValide = [
        'Standard Crossfit', 
        'Student Crossfit', 
        'Open Gym', 
        'Premium All Access', 
        'Full Time', 
        'Weekend Only'
    ];

    if (!tipuriValide.includes(tip)) {
        alert(`Eroare: Tipul abonamentului trebuie să fie unul dintre: \n${tipuriValide.join(', ')}`);
        return;
    }

    if (isNaN(pret) || pret < 0) {
        alert('Eroare: Prețul abonamentului trebuie să fie un număr pozitiv (>= 0).');
        return;
    }

    const abonamentData = {
        tip: tip,
        pret: pret
    };

    let url = '/api/abonamente/add';

    if (currentEditID !== 0) {
        abonamentData.id = currentEditID;
        url = '/api/abonamente/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(abonamentData)
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la salvarea abonamentului'); }
            return response.json();
        })
        .then(data => {
            alert(data.mesaj);
            resetFormular();
            loadAbonamente();
        })
        .catch(error => alert('Eroare formular: ' + error.message));
}

function handleEditClick(id) {
    fetch(`/api/abonament?id=${id}`) // API nou pentru un singur abonament
        .then(response => response.json())
        .then(abonament => {
            document.getElementById('tip-abonament').value = abonament.tip;
            document.getElementById('pret-abonament').value = abonament.pret;

            currentEditID = abonament.id;
            document.querySelector('#form-add-abonament button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor abonamentului:', error));
}

function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi abonamentul cu ID-ul ${id}?`)) {
        return;
    }

    fetch('/api/abonamente/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la ștergerea abonamentului'); }
            return response.json();
        })
        .then(data => {
            alert(data.mesaj);
            loadAbonamente();
        })
        .catch(error => alert('Eroare: ' + error.message));
}

function resetFormular() {
    document.getElementById('form-add-abonament').reset();
    currentEditID = 0;
    document.querySelector('#form-add-abonament button[type="submit"]').textContent = 'Adaugă Abonament';
}