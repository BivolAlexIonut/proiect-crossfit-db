let currentEditID = 0;
let claseList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    // Încărcăm tot ce e necesar pentru pagină
    loadClase();
    loadAntrenoriDropdown();
    loadCategoriiDropdown();

    document.getElementById('form-add-clasa').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-clase').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            handleDelete(event.target.getAttribute('data-id'));
        }
        if (event.target.classList.contains('btn-edit')) {
            handleEditClick(event.target.getAttribute('data-id'));
        }
    });
});

// Încarcă tabelul principal cu clase
function loadClase() {
    fetch('/api/clase')
        .then(response => response.json())
        .then(data => {
            claseList = data; // Store data
            renderClase();   // Render
        })
        .catch(error => console.error('Eroare la preluarea claselor:', error));
}

function renderClase() {
    const tbody = document.getElementById('lista-clase');
    tbody.innerHTML = '';
    if (!claseList) claseList = [];
    claseList.forEach(clasa => {
        const tr = document.createElement('tr');
        // Formatăm dataOra să fie mai lizibilă
        const dataOra = new Date(clasa.dataOra).toLocaleString('ro-RO');

        tr.innerHTML = `
            <td>${clasa.id}</td>
            <td>${clasa.numeWOD}</td>
            <td>${clasa.descriereWOD || '-'}</td>
            <td>${dataOra}</td>
            <td>${clasa.numeAntrenor || '-'}</td> <td>${clasa.numeCategorie || '-'}</td> <td>
                <button class="btn-edit" data-id="${clasa.id}">Editează</button>
                <button class="btn-delete" data-id="${clasa.id}">Șterge</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function sortClase(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    claseList.sort((a, b) => {
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

    renderClase();
}

// Încarcă antrenorii în dropdown
function loadAntrenoriDropdown() {
    fetch('/api/antrenori')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-antrenor');
            data.forEach(antrenor => {
                const option = document.createElement('option');
                option.value = antrenor.id;
                option.textContent = `${antrenor.nume} ${antrenor.prenume}`;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare la preluarea antrenorilor:', error));
}

// Încarcă categoriile în dropdown
function loadCategoriiDropdown() {
    fetch('/api/tipuri-antrenament')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-categorie');
            data.forEach(categorie => {
                const option = document.createElement('option');
                option.value = categorie.id;
                option.textContent = categorie.nume;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare la preluarea categoriilor:', error));
}

// Gestionează Adăugare / Editare
function handleFormSubmit(event) {
    event.preventDefault();

    const clasaData = {
        numeWOD: document.getElementById('nume-wod').value,
        descriereWOD: document.getElementById('descriere-wod').value,
        dataOra: document.getElementById('data-ora').value,
        antrenorID: parseInt(document.getElementById('select-antrenor').value, 10),
        tipAntrenamentID: parseInt(document.getElementById('select-categorie').value, 10)
    };

    let url = '/api/clase/add';
    if (currentEditID !== 0) {
        clasaData.id = currentEditID;
        url = '/api/clase/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(clasaData)
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la salvarea clasei'); } return response.json(); })
        .then(data => { console.log(data.mesaj); resetFormular(); loadClase(); })
        .catch(error => console.error('Eroare formular:', error));
}

// Populează formularul la click pe "Editează"
function handleEditClick(id) {
    fetch(`/api/clasa?id=${id}`)
        .then(response => response.json())
        .then(clasa => {
            // Oracle trimite data în format ISO 8601, tăiem ":ss.sssZ" de la final
            const dataOraLocal = clasa.dataOra.substring(0, 16);

            document.getElementById('nume-wod').value = clasa.numeWOD;
            document.getElementById('descriere-wod').value = clasa.descriereWOD;
            document.getElementById('data-ora').value = dataOraLocal;
            document.getElementById('select-antrenor').value = clasa.antrenorID;
            document.getElementById('select-categorie').value = clasa.tipAntrenamentID;

            currentEditID = clasa.id;
            document.querySelector('#form-add-clasa button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor clasei:', error));
}

// Șterge o clasă
function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi clasa cu ID-ul ${id}? Toate înscrierile asociate vor fi șterse!`)) { return; }

    fetch('/api/clase/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la ștergerea clasei'); } return response.json(); })
        .then(data => { console.log(data.mesaj); loadClase(); })
        .catch(error => console.error('Eroare la ștergere:', error));
}

// Resetează formularul
function resetFormular() {
    document.getElementById('form-add-clasa').reset();
    currentEditID = 0;
    document.querySelector('#form-add-clasa button[type="submit"]').textContent = 'Adaugă Clasă';
}