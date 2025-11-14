let currentEditID = 0;

// Mapare pentru a converti numărul zilei în text
const zileSaptamana = {
    1: 'Luni',
    2: 'Marți',
    3: 'Miercuri',
    4: 'Joi',
    5: 'Vineri',
    6: 'Sâmbătă',
    7: 'Duminică'
};

window.addEventListener('DOMContentLoaded', () => {
    loadOrar();
    loadAntrenoriDropdown();
    loadCategoriiDropdown();

    document.getElementById('form-add-orar').addEventListener('submit', handleFormSubmit);
    document.getElementById('btn-generate-orar').addEventListener('click', handleGenerateOrar);
    document.getElementById('lista-orar').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            handleDelete(event.target.getAttribute('data-id'));
        }
        if (event.target.classList.contains('btn-edit')) {
            handleEditClick(event.target.getAttribute('data-id'));
        }
    });
});

// Încarcă tabelul cu regulile de orar
function loadOrar() {
    fetch('/api/orar')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('lista-orar');
            tbody.innerHTML = '';
            if (!data) data = [];
            data.forEach(regula => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${regula.id}</td>
                    <td>${zileSaptamana[regula.ziuaSaptamanii] || 'N/A'}</td>
                    <td>${regula.ora}</td>
                    <td>${regula.numeWODTemplate}</td>
                    <td>${regula.numeAntrenor || 'N/A'}</td>
                    <td>${regula.numeCategorie || 'N/A'}</td>
                    <td>
                        <button class="btn-edit" data-id="${regula.id}">Editează</button>
                        <button class="btn-delete" data-id="${regula.id}">Șterge</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea orarului:', error));
}

// Încarcă antrenorii în dropdown
function loadAntrenoriDropdown() {
    fetch('/api/antrenori')
        .then(response => response.json())
        .then(data => {
            const select = document.getElementById('select-antrenor-orar');
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
            const select = document.getElementById('select-categorie-orar');
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

    const orarData = {
        ziuaSaptamanii: parseInt(document.getElementById('select-zi').value, 10),
        ora: document.getElementById('ora').value,
        numeWODTemplate: document.getElementById('nume-wod-template').value,
        antrenorID: parseInt(document.getElementById('select-antrenor-orar').value, 10),
        tipAntrenamentID: parseInt(document.getElementById('select-categorie-orar').value, 10)
    };

    let url = '/api/orar/add';
    if (currentEditID !== 0) {
        orarData.id = currentEditID;
        url = '/api/orar/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(orarData)
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la salvarea regulii'); } return response.json(); })
        .then(data => { console.log(data.mesaj); resetFormular(); loadOrar(); })
        .catch(error => console.error('Eroare formular:', error));
}

// Populează formularul la click pe "Editează"
function handleEditClick(id) {
    fetch(`/api/orar/single?id=${id}`)
        .then(response => response.json())
        .then(regula => {
            document.getElementById('select-zi').value = regula.ziuaSaptamanii;
            document.getElementById('ora').value = regula.ora;
            document.getElementById('nume-wod-template').value = regula.numeWODTemplate;
            document.getElementById('select-antrenor-orar').value = regula.antrenorID;
            document.getElementById('select-categorie-orar').value = regula.tipAntrenamentID;

            currentEditID = regula.id;
            document.querySelector('#form-add-orar button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor regulii:', error));
}

// Șterge o regulă
function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi această regulă de orar (ID: ${id})?`)) { return; }

    fetch('/api/orar/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la ștergerea regulii'); } return response.json(); })
        .then(data => { console.log(data.mesaj); loadOrar(); })
        .catch(error => console.error('Eroare la ștergere:', error));
}

// Resetează formularul
function resetFormular() {
    document.getElementById('form-add-orar').reset();
    currentEditID = 0;
    document.querySelector('#form-add-orar button[type="submit"]').textContent = 'Adaugă Regulă';
}

// --- FUNCȚIE NOUĂ PENTRU GENERAREA ORARULUI ---
function handleGenerateOrar() {
    const statusEl = document.getElementById('generate-status');
    if (!confirm('Ești sigur că vrei să generezi clasele pentru săptămâna viitoare?')) {
        return;
    }

    statusEl.textContent = 'Se generează... Vă rugăm așteptați...';
    statusEl.style.color = '#007bff';

    fetch('/api/orar/generate', {
        method: 'POST'
    })
        .then(response => {
            if (!response.ok) {
                // Dacă eșuează, încercăm să citim textul erorii
                return response.text().then(text => { throw new Error(text) });
            }
            return response.json();
        })
        .then(data => {
            statusEl.textContent = `Operațiune reușită! Au fost generate ${data.claseGenerate} clase.`;
            statusEl.style.color = 'green';
            // Opțional: Redirecționează utilizatorul să vadă rezultatele
            setTimeout(() => {
                window.location.href = '/clase';
            }, 2000);
        })
        .catch(error => {
            console.error('Eroare la generare:', error);
            statusEl.textContent = `Eroare: ${error.message}`;
            statusEl.style.color = 'red';
        });
}