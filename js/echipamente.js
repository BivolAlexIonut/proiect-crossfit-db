let currentEditID = 0;
let echipamenteList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    loadEchipamente();

    document.getElementById('form-add-echipament').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-echipamente').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const echipamentID = event.target.getAttribute('data-id');
            handleDelete(echipamentID);
        }
        if (event.target.classList.contains('btn-edit')) {
            const echipamentID = event.target.getAttribute('data-id');
            handleEditClick(echipamentID);
        }
    });
});

function loadEchipamente() {
    fetch('/api/echipamente')
        .then(response => response.json())
        .then(data => {
            echipamenteList = data;
            renderEchipamente();
        })
        .catch(error => alert('Eroare la preluarea echipamentelor: ' + error.message));
}

function renderEchipamente() {
    const tbody = document.getElementById('lista-echipamente');
    tbody.innerHTML = '';
    if (!echipamenteList) echipamenteList = [];
    echipamenteList.forEach(echipament => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${echipament.id}</td>
            <td>${echipament.nume}</td>
            <td>${echipament.cantitate}</td>
            <td>
                <button class="btn-edit" data-id="${echipament.id}">Editează</button>
                <button class="btn-delete" data-id="${echipament.id}">Șterge</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function sortEchipamente(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    echipamenteList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderEchipamente();
}

function handleFormSubmit(event) {
    event.preventDefault();

    const cantitate = parseInt(document.getElementById('cantitate-echipament').value, 10);

    if (isNaN(cantitate) || cantitate < 0) {
        alert('Eroare: Cantitatea totală trebuie să fie un număr pozitiv (>= 0).');
        return;
    }

    const echipamentData = {
        nume: document.getElementById('nume-echipament').value,
        cantitate: cantitate
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
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la salvarea echipamentului'); }
            return response.json();
        })
        .then(data => {
            alert(data.mesaj);
            resetFormular();
            loadEchipamente();
        })
        .catch(error => alert('Eroare formular: ' + error.message));
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
        .catch(error => alert('Eroare: ' + error.message));
}

function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi echipamentul cu ID-ul ${id}?`)) {
        return;
    }

    fetch('/api/echipamente/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la ștergerea echipamentului'); }
            return response.json();
        })
        .then(data => {
            alert(data.mesaj);
            loadEchipamente();
        })
        .catch(error => alert('Eroare: ' + error.message));
}

function resetFormular() {
    document.getElementById('form-add-echipament').reset();
    currentEditID = 0;
    document.querySelector('#form-add-echipament button[type="submit"]').textContent = 'Adaugă Echipament';
}