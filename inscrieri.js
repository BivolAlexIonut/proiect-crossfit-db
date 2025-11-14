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
            const tbody = document.getElementById('lista-inscrieri');
            tbody.innerHTML = '';
            if (!data) data = [];
            data.forEach(inscriere => {
                const tr = document.createElement('tr');
                const dataOra = new Date(inscriere.dataOra).toLocaleString('ro-RO');

                tr.innerHTML = `
                    <td>${inscriere.numeMembru}</td>
                    <td>${inscriere.numeWOD}</td>
                    <td>${dataOra}</td>
                    <td>${inscriere.numeAntrenor}</td>
                    <td>
                        <button class="btn-delete" 
                                data-membru-id="${inscriere.membruID}" 
                                data-clasa-id="${inscriere.clasaID}">
                            Anulează
                        </button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea înscrierilor:', error));
}

function loadMembriDropdown() {
    fetch('/api/membri')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-membru');
            data.forEach(membru => {
                const option = document.createElement('option');
                option.value = membru.id;
                option.textContent = `${membru.nume} ${membru.prenume} (ID: ${membru.id})`;
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
                const dataOra = new Date(clasa.dataOra).toLocaleString('ro-RO');
                const option = document.createElement('option');
                option.value = clasa.id;
                option.textContent = `${clasa.numeWOD} @ ${dataOra}`;
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
            if (!response.ok) { throw new Error('Eroare la salvarea înscrierii. Membrul este deja înscris?'); }
            return response.json();
        })
        .then(data => {
            console.log(data.mesaj);
            document.getElementById('form-add-inscriere').reset();
            loadInscrieri();
        })
        .catch(error => alert(error.message));
}

function handleDelete(membruID, clasaID) {
    if (!confirm(`Ești sigur că vrei să anulezi această înscriere?`)) { return; }

    fetch('/api/inscrieri/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            membruID: parseInt(membruID, 10),
            clasaID: parseInt(clasaID, 10)
        })
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la anularea înscrierii'); } return response.json(); })
        .then(data => { console.log(data.mesaj); loadInscrieri(); })
        .catch(error => console.error('Eroare la ștergere:', error));
}