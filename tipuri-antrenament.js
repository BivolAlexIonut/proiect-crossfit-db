let currentEditID = 0;

window.addEventListener('DOMContentLoaded', () => {
    loadTipuri();
    document.getElementById('form-add-tip').addEventListener('submit', handleFormSubmit);
    document.getElementById('lista-tipuri').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            handleDelete(event.target.getAttribute('data-id'));
        }
        if (event.target.classList.contains('btn-edit')) {
            handleEditClick(event.target.getAttribute('data-id'));
        }
    });
});

function loadTipuri() {
    fetch('/api/tipuri-antrenament')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('lista-tipuri');
            tbody.innerHTML = '';
            data.forEach(tip => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${tip.id}</td>
                    <td>${tip.nume}</td>
                    <td>${tip.descriere}</td>
                    <td>
                        <button class="btn-edit" data-id="${tip.id}">Editează</button>
                        <button class="btn-delete" data-id="${tip.id}">Șterge</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea tipurilor:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();
    const tipData = {
        nume: document.getElementById('nume-tip').value,
        descriere: document.getElementById('descriere-tip').value
    };

    let url = '/api/tipuri-antrenament/add';
    if (currentEditID !== 0) {
        tipData.id = currentEditID;
        url = '/api/tipuri-antrenament/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(tipData)
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la salvarea categoriei'); } return response.json(); })
        .then(data => { console.log(data.mesaj); resetFormular(); loadTipuri(); })
        .catch(error => console.error('Eroare formular:', error));
}

function handleEditClick(id) {
    fetch(`/api/tip-antrenament?id=${id}`)
        .then(response => response.json())
        .then(tip => {
            document.getElementById('nume-tip').value = tip.nume;
            document.getElementById('descriere-tip').value = tip.descriere;
            currentEditID = tip.id;
            document.querySelector('#form-add-tip button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor:', error));
}

function handleDelete(id) {
    // Am schimbat textul aici
    if (!confirm(`Ești sigur că vrei să ștergi această categorie (ID: ${id})?`)) { return; }
    fetch('/api/tipuri-antrenament/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la ștergerea categoriei'); } return response.json(); })
        .then(data => { console.log(data.mesaj); loadTipuri(); })
        .catch(error => console.error('Eroare la ștergere:', error));
}

function resetFormular() {
    document.getElementById('form-add-tip').reset();
    currentEditID = 0;
    // Am schimbat textul aici
    document.querySelector('#form-add-tip button[type="submit"]').textContent = 'Adaugă Categorie';
}