/**
 * =============================================================================
 * SCRIPT MANAGEMENT MEMBRI
 * =============================================================================
 * Acest script gestionează operațiunile CRUD pentru entitatea Membri.
 * Interacționează cu API-ul backend definit în main.go.
 */

let currentEditID = 0; // ID-ul membrului aflat în editare (0 = mod adăugare)
let membriList = [];   // Cache local pentru sortare rapidă
let sortDirection = 1; // Direcția de sortare: 1 (asc), -1 (desc)
let lastSortColumn = '';

// Inițializare la încărcarea paginii
document.addEventListener('DOMContentLoaded', () => {
    loadMembri();
    loadAbonamente();
    
    // Atașare event listener pentru formular
    document.getElementById('form-add-membru').addEventListener('submit', handleFormSubmit);

    // Delegare evenimente pentru butoane (Editare/Ștergere)
    document.getElementById('lista-membri').addEventListener('click', (event) => {
        const target = event.target;
        const id = target.getAttribute('data-id');

        if (target.classList.contains('btn-delete')) {
            handleDeleteMembru(id);
        } else if (target.classList.contains('btn-edit')) {
            handleEditClick(id);
        }
    });
});

/**
 * Încarcă lista de membri de la server.
 */
async function loadMembri() {
    try {
        const response = await fetch('/api/membri');
        if (!response.ok) throw new Error('Eroare rețea');
        
        const data = await response.json();
        membriList = data || []; // Asigurăm că e array
        renderMembri();
    } catch (error) {
        alert('Eroare la preluarea membrilor: ' + error.message);
    }
}

/**
 * Randează tabelul HTML pe baza datelor din membriList.
 */
function renderMembri() {
    const tbody = document.getElementById('lista-membri');
    if (!tbody) return;
    tbody.innerHTML = ''; // Curățare tabel

    if (membriList.length === 0) {
        tbody.innerHTML = '<tr><td colspan="6">Nu există membri înregistrați.</td></tr>';
        return;
    }

    membriList.forEach(membru => {
        const tr = document.createElement('tr');
        // Construire rând cu atribute data-label pentru responsive design
        tr.innerHTML = `
            <td data-label="ID Membru">${membru.id}</td>
            <td data-label="Nume">${membru.nume}</td>
            <td data-label="Prenume">${membru.prenume}</td>
            <td data-label="Email">${membru.email}</td>
            <td data-label="Abonament">${membru.tipAbonament || '-'}</td>
            <td data-label="Acțiuni">
                <button class="btn-edit" data-id="${membru.id}">Editează</button>
                <button class="btn-delete" data-id="${membru.id}">Șterge</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

/**
 * Încarcă lista de abonamente pentru dropdown-ul din formular.
 */
async function loadAbonamente() {
    try {
        const response = await fetch('/api/abonamente');
        const data = await response.json();
        
        const select = document.getElementById('select-abonament');
        if (!select) return;
        select.innerHTML = '<option value="">Alege un abonament...</option>';
        
        if (data) {
            data.forEach(ab => {
                const option = document.createElement('option');
                option.value = ab.id;
                option.textContent = `${ab.tip} - ${ab.pret} RON`;
                select.appendChild(option);
            });
        }
    } catch (error) {
        alert('Eroare la preluarea abonamentelor: ' + error.message);
    }
}

/**
 * Gestionează trimiterea formularului (Adăugare sau Editare).
 */
async function handleFormSubmit(event) {
    event.preventDefault();

    const abonamentValue = document.getElementById('select-abonament').value;
    if (!abonamentValue) {
        alert("Te rugăm să selectezi un abonament!");
        return;
    }

    const membruData = {
        nume: document.getElementById('nume').value,
        prenume: document.getElementById('prenume').value,
        email: document.getElementById('email').value,
        abonamentID: parseInt(abonamentValue, 10)
    };

    // Determinăm dacă este adăugare sau editare
    const url = currentEditID === 0 ? '/api/membri/add' : '/api/membri/update';
    if (currentEditID !== 0) {
        membruData.id = currentEditID;
    }

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(membruData)
        });

        if (!response.ok) throw new Error('Operațiune eșuată');
        
        const result = await response.json();
        alert(result.mesaj);
        
        resetFormular();
        loadMembri();
    } catch (error) {
        alert('Eroare formular: ' + error.message);
    }
}

/**
 * Pregătește formularul pentru editarea unui membru existent.
 */
async function handleEditClick(id) {
    try {
        const response = await fetch(`/api/membru?id=${id}`);
        if (!response.ok) throw new Error('Nu s-a găsit membrul');
        
        const membru = await response.json();

        // Populare câmpuri
        document.getElementById('nume').value = membru.nume;
        document.getElementById('prenume').value = membru.prenume;
        document.getElementById('email').value = membru.email;
        document.getElementById('select-abonament').value = membru.abonamentID;

        // Setare stare editare
        currentEditID = membru.id;
        document.querySelector('#form-add-membru button[type="submit"]').textContent = 'Salvează Modificările';
        
        // Scroll la formular
        window.scrollTo({ top: 0, behavior: 'smooth' });
    } catch (error) {
        alert('Eroare: ' + error.message);
    }
}

/**
 * Șterge un membru (cu confirmare și logică de cascadă).
 */
async function handleDeleteMembru(id) {
    const confirmation = confirm(
        `ATENȚIE! Ștergerea membrului ID ${id} va șterge automat și:\n` +
        `- Istoricul achizițiilor\n` +
        `- Înscrierile la clase\n` +
        `- Mentoratele active\n\n` +
        `Continui?`
    );

    if (!confirmation) return;

    try {
        const response = await fetch('/api/membri/delete', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ id: parseInt(id, 10) })
        });

        if (!response.ok) throw new Error('Ștergere eșuată');
        
        const result = await response.json();
        alert(result.mesaj);
        loadMembri();
    } catch (error) {
        alert('Eroare: ' + error.message);
    }
}

/**
 * Resetează formularul la starea inițială.
 */
function resetFormular() {
    document.getElementById('form-add-membru').reset();
    currentEditID = 0;
    document.querySelector('#form-add-membru button[type="submit"]').textContent = 'Adaugă Membru';
}

/**
 * Sortează tabelul membrilor.
 */
function sortMembri(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    membriList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderMembri();
}
