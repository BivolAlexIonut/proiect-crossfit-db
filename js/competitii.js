let competitiiList = [];
let participariList = [];
let sortDirection = 1;
let lastSortColumn = '';

document.addEventListener('DOMContentLoaded', () => {
    loadCompetitii();
    loadParticipari(); // Încărcăm tabelul de participări
    loadSelectsParticipare(); // Încărcăm dropdown-urile pentru formularul de înscriere

    document.getElementById('form-add-competitie').addEventListener('submit', handleAddCompetitie);
    document.getElementById('form-add-participare').addEventListener('submit', handleAddParticipare);
});

// --- COMPETITII ---
function loadCompetitii() {
    fetch('/api/competitii')
        .then(res => res.json())
        .then(data => {
            competitiiList = data || [];
            renderCompetitii();
        })
        .catch(err => console.error(err));
}

function renderCompetitii() {
    const tbody = document.getElementById('lista-competitii');
    tbody.innerHTML = '';
    competitiiList.forEach(c => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${c.id}</td>
            <td>${c.nume}</td>
            <td>${c.data}</td>
            <td>${c.locatie}</td>
            <td>${c.taxa} RON</td>
            <td><button onclick="deleteCompetitie(${c.id})">Șterge</button></td>
        `;
        tbody.appendChild(tr);
    });
}

function sortCompetitii(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    competitiiList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        if (valA === null || valA === undefined) valA = '';
        if (valB === null || valB === undefined) valB = '';

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderCompetitii();
}

function handleAddCompetitie(e) {
    e.preventDefault();
    const data = {
        nume: document.getElementById('nume-competitie').value,
        data: document.getElementById('data-competitie').value,
        locatie: document.getElementById('locatie-competitie').value,
        taxa: parseFloat(document.getElementById('taxa-competitie').value)
    };

    fetch('/api/competitii/add', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
    .then(res => {
        if(!res.ok) throw new Error('Eroare adăugare');
        loadCompetitii();
        loadSelectsParticipare(); // Reîncărcăm și dropdown-ul de competiții
        e.target.reset();
    })
    .catch(err => alert('Eroare: ' + err));
}

function deleteCompetitie(id) {
    if(!confirm('Ștergi competiția?')) return;
    fetch('/api/competitii/delete', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({id: id})
    })
    .then(() => {
        loadCompetitii();
        loadSelectsParticipare();
    })
    .catch(err => console.error(err));
}

// --- PARTICIPARI ---
function loadParticipari() {
    fetch('/api/competitii/participari')
        .then(res => res.json())
        .then(data => {
            participariList = data || [];
            renderParticipari();
        })
        .catch(err => console.error(err));
}

function renderParticipari() {
    const tbody = document.getElementById('lista-participari');
    tbody.innerHTML = '';
    participariList.forEach(p => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${p.numeCompetitie}</td>
            <td>${p.numeMembru}</td>
            <td>${p.loculObtinut > 0 ? p.loculObtinut : '-'}</td>
        `;
        tbody.appendChild(tr);
    });
}

// Sortare pentru Participări (opțional, dar bun pentru consistență)
// Presupunem că adăugăm onclick și în HTML-ul pentru participări dacă vrem sortare acolo
function sortParticipari(column) {
     // Implementare similară dacă se cere sortare și pe tabelul secundar
}

function loadSelectsParticipare() {
    // Competitii
    fetch('/api/competitii')
        .then(res => res.json())
        .then(data => {
            const sel = document.getElementById('select-competitie');
            sel.innerHTML = '<option value="">Alege Competiția...</option>';
            if(data) data.forEach(c => {
                const opt = document.createElement('option');
                opt.value = c.id;
                opt.textContent = `${c.nume} (${c.data})`;
                sel.appendChild(opt);
            });
        });

    // Membri
    fetch('/api/membri')
        .then(res => res.json())
        .then(data => {
            const sel = document.getElementById('select-membru-comp');
            sel.innerHTML = '<option value="">Alege Membrul...</option>';
            if(data) data.forEach(m => {
                const opt = document.createElement('option');
                opt.value = m.id;
                opt.textContent = `${m.nume} ${m.prenume}`;
                sel.appendChild(opt);
            });
        });
}

function handleAddParticipare(e) {
    e.preventDefault();
    const data = {
        competitieID: parseInt(document.getElementById('select-competitie').value),
        membruID: parseInt(document.getElementById('select-membru-comp').value)
    };

    fetch('/api/competitii/participari/add', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
    .then(res => {
        if(!res.ok) throw new Error('Eroare înscriere');
        alert('Membru înscris!');
        loadParticipari();
    })
    .catch(err => alert('Eroare: ' + err));
}