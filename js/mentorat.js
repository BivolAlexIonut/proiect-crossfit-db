let mentoratList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    loadMentorat();
    loadAntrenoriDropdown();
    loadMembriDropdown();

    document.getElementById('form-add-mentorat').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-mentorat').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const antrenorID = event.target.getAttribute('data-antrenor-id');
            const membruID = event.target.getAttribute('data-membru-id');
            handleDelete(antrenorID, membruID);
        }
    });
});

function loadMentorat() {
    fetch('/api/mentorat')
        .then(response => response.json())
        .then(data => {
            mentoratList = data;
            renderMentorat();
        })
        .catch(error => console.error('Eroare la preluarea mentoratelor:', error));
}

function renderMentorat() {
    const tbody = document.getElementById('lista-mentorat');
    tbody.innerHTML = '';
    if (!mentoratList) mentoratList = [];
    mentoratList.forEach(item => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${item.numeAntrenor}</td>
            <td>${item.numeMembru}</td>
            <td>
                <button class="btn-delete" data-antrenor-id="${item.antrenorID}" data-membru-id="${item.membruID}">Șterge</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function sortMentorat(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    mentoratList.sort((a, b) => {
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

    renderMentorat();
}

function loadAntrenoriDropdown() {
    fetch('/api/antrenori')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-antrenor');
            data.forEach(a => {
                const option = document.createElement('option');
                option.value = a.id;
                option.textContent = `${a.nume} ${a.prenume}`;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare load antrenori:', error));
}

function loadMembriDropdown() {
    fetch('/api/membri')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-membru');
            data.forEach(m => {
                const option = document.createElement('option');
                option.value = m.id;
                option.textContent = `${m.nume} ${m.prenume}`;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare load membri:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();
    const data = {
        antrenorID: parseInt(document.getElementById('select-antrenor').value),
        membruID: parseInt(document.getElementById('select-membru').value)
    };

    fetch('/api/mentorat/add', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) throw new Error('Eroare la adăugare');
        return response.json();
    })
    .then(json => {
        console.log(json.mesaj);
        loadMentorat();
    })
    .catch(err => console.error(err));
}

function handleDelete(antrenorID, membruID) {
    if (!confirm('Ștergi acest mentorat?')) return;

    fetch('/api/mentorat/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ antrenorID: parseInt(antrenorID), membruID: parseInt(membruID) })
    })
    .then(response => {
        if (!response.ok) throw new Error('Eroare la ștergere');
        loadMentorat();
    })
    .catch(err => console.error(err));
}