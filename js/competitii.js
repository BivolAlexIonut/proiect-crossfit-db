let competitiiList = [];
let participariList = [];
let sortDirection = 1;
let lastSortColumn = '';

document.addEventListener('DOMContentLoaded', () => {
    loadCompetitii();
    loadParticipari();
    loadSelectsParticipare();

    const formAdd = document.getElementById('form-add-competitie');
    if (formAdd) {
        formAdd.addEventListener('submit', handleAddCompetitie);
    }

    const formInscriere = document.getElementById('form-inscriere-comp');
    if (formInscriere) {
        formInscriere.addEventListener('submit', handleAddParticipare);
    }
});

// --- COMPETITII ---
function loadCompetitii() {
    fetch('/api/competitii')
        .then(res => res.json())
        .then(data => {
            competitiiList = data || [];
            renderCompetitii();
        })
        .catch(err => alert('Eroare la încărcarea competițiilor: ' + err.message));
}

function renderCompetitii() {
    const tbody = document.getElementById('lista-competitii');
    if (!tbody) return;
    tbody.innerHTML = '';
    
    if (competitiiList.length === 0) {
        tbody.innerHTML = '<tr><td colspan="6">Nu există competiții.</td></tr>';
        return;
    }

    competitiiList.forEach(c => {
        const tr = document.createElement('tr');
        const dataFormatata = new Date(c.data).toLocaleDateString('ro-RO');
        tr.innerHTML = `
            <td data-label="ID">${c.id}</td>
            <td data-label="Nume">${c.nume}</td>
            <td data-label="Data">${dataFormatata}</td>
            <td data-label="Locație">${c.locatie}</td>
            <td data-label="Taxă">${c.taxa.toFixed(2)} RON</td>
            <td data-label="Acțiuni"><button class="btn-delete" onclick="deleteCompetitie(${c.id})">Șterge</button></td>
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
        nume: document.getElementById('nume').value,
        data: document.getElementById('data').value,
        locatie: document.getElementById('locatie').value,
        taxa: parseFloat(document.getElementById('taxa').value)
    };

    fetch('/api/competitii/add', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
    .then(async res => {
        if (!res.ok) {
            const text = await res.text();
            throw new Error(text || 'Eroare la adăugare');
        }
        return res.json();
    })
    .then(resp => {
        alert(resp.mesaj);
        loadCompetitii();
        loadSelectsParticipare();
        document.getElementById('form-add-competitie').reset();
    })
    .catch(err => alert('Eroare: ' + err.message));
}

function deleteCompetitie(id) {
    if (!confirm('Sigur ștergi competiția?')) return;
    
    fetch('/api/competitii/delete', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({id: id})
    })
    .then(async res => {
        if (!res.ok) {
            const text = await res.text();
            throw new Error(text || 'Eroare la ștergere');
        }
        return res.json();
    })
    .then(resp => {
        alert(resp.mesaj);
        loadCompetitii();
        loadSelectsParticipare();
    })
    .catch(err => alert('Eroare: ' + err.message));
}

// --- PARTICIPARI ---
function loadParticipari() {
    fetch('/api/competitii/participari')
        .then(res => res.json())
        .then(data => {
            participariList = data || [];
            renderParticipari();
        })
        .catch(err => console.error(err)); // Aici lasam console.error sau alert, e ok
}

function renderParticipari() {
    const tbody = document.getElementById('lista-participari');
    if (!tbody) return;
    tbody.innerHTML = '';
    
    if (participariList.length === 0) {
        tbody.innerHTML = '<tr><td colspan="3">Nu există participanți înscriși.</td></tr>';
        return;
    }

    participariList.forEach(p => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td data-label="Eveniment">${p.numeCompetitie}</td>
            <td data-label="Membru">${p.numeMembru}</td>
            <td data-label="Loc Obținut">${p.loculObtinut > 0 ? p.loculObtinut : '-'}</td>
        `;
        tbody.appendChild(tr);
    });
}

function loadSelectsParticipare() {
    // Competitii
    fetch('/api/competitii')
        .then(res => res.json())
        .then(data => {
            const sel = document.getElementById('select-competitie');
            if (!sel) return;
            sel.innerHTML = '<option value="">Alege Competiția...</option>';
            if(data) data.forEach(c => {
                const opt = document.createElement('option');
                opt.value = c.id;
                opt.textContent = `${c.nume} (${new Date(c.data).toLocaleDateString('ro-RO')})`;
                sel.appendChild(opt);
            });
        });

    // Membri
    fetch('/api/membri')
        .then(res => res.json())
        .then(data => {
            const sel = document.getElementById('select-membru'); // ID corectat conform HTML
            if (!sel) return;
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
        membruID: parseInt(document.getElementById('select-membru').value)
    };

    fetch('/api/competitii/participari/add', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
    .then(async res => {
        if (!res.ok) {
            const text = await res.text();
            throw new Error(text || 'Eroare înscriere');
        }
        return res.json();
    })
    .then(resp => {
        alert(resp.mesaj);
        loadParticipari();
    })
    .catch(err => alert('Eroare: ' + err.message));
}