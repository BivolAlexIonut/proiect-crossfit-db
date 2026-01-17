let inscrieriList = [];
let sortDirection = 1;
let lastSortColumn = '';

window.addEventListener('DOMContentLoaded', () => {
    loadInscrieri();
    loadMembriDropdown();
    loadClaseDropdown();

    document.getElementById('form-add-inscriere').addEventListener('submit', handleFormSubmit);

    document.getElementById('lista-inscrieri').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            const membruID = event.target.getAttribute('data-membru-id');
            const clasaID = event.target.getAttribute('data-clasa-id');
            handleDelete(membruID, clasaID);
        }
    });
});

function loadInscrieri() {
    fetch('/api/inscrieri')
        .then(response => response.json())
        .then(data => {
            inscrieriList = data;
            renderInscrieri();
        })
        .catch(error => console.error('Eroare la preluarea înscrierilor:', error));
}

function renderInscrieri() {
    const tbody = document.getElementById('lista-inscrieri');
    tbody.innerHTML = '';
    if (!inscrieriList) inscrieriList = [];
    inscrieriList.forEach(inscriere => {
        const tr = document.createElement('tr');
        
        // Formatare dată
        const dataFormata = new Date(inscriere.dataOra).toLocaleString('ro-RO');

        tr.innerHTML = `
            <td>${inscriere.numeMembru}</td>
            <td>${inscriere.numeWOD}</td>
            <td>${dataFormata}</td>
            <td>${inscriere.numeAntrenor || '-'}</td>
            <td>
                <button class="btn-delete" data-membru-id="${inscriere.membruID}" data-clasa-id="${inscriere.clasaID}">Anulează</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function sortInscrieri(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    inscrieriList.sort((a, b) => {
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

    renderInscrieri();
}


function loadMembriDropdown() {
    fetch('/api/membri')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-membru');
            data.forEach(membru => {
                const option = document.createElement('option');
                option.value = membru.id;
                option.textContent = `${membru.nume} ${membru.prenume}`;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare la preluarea membrilor:', error));
}

function loadClaseDropdown() {
    fetch('/api/clase')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-clasa');
            data.forEach(clasa => {
                const option = document.createElement('option');
                option.value = clasa.id;
                // Formatare simplă pentru dropdown
                const dataLocala = new Date(clasa.dataOra).toLocaleString('ro-RO');
                option.textContent = `${clasa.numeWOD} (${dataLocala})`;
                select.appendChild(option);
            });
        })
        .catch(error => console.error('Eroare la preluarea claselor:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();

    const inscriereData = {
        membruID: parseInt(document.getElementById('select-membru').value, 10),
        clasaID: parseInt(document.getElementById('select-clasa').value, 10)
    };

    fetch('/api/inscrieri/add', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(inscriereData)
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text); });
            }
            return response.json();
        })
        .then(data => {
            alert(data.mesaj);
            document.getElementById('form-add-inscriere').reset();
            loadInscrieri();
        })
        .catch(error => {
            console.error('Eroare formular:', error);
            alert(error.message); // Afișăm eroarea venită din backend (ex: limita de ședințe)
        });
}

function handleDelete(membruID, clasaID) {
    if (!confirm(`Sigur vrei să anulezi înscrierea?`)) {
        return;
    }

    fetch('/api/inscrieri/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ membruID: parseInt(membruID, 10), clasaID: parseInt(clasaID, 10) })
    })
        .then(response => {
            if (!response.ok) { throw new Error('Eroare la anularea înscrierii'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            loadInscrieri();
        })
        .catch(error => console.error('Eroare la ștergere:', error));
}