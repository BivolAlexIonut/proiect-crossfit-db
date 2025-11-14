let currentEditID = 0;

window.addEventListener('DOMContentLoaded', () => {
    loadProduse();
    document.getElementById('form-add-produs').addEventListener('submit', handleFormSubmit);
    document.getElementById('lista-produse').addEventListener('click', (event) => {
        if (event.target.classList.contains('btn-delete')) {
            handleDelete(event.target.getAttribute('data-id'));
        }
        if (event.target.classList.contains('btn-edit')) {
            handleEditClick(event.target.getAttribute('data-id'));
        }
    });
});

function loadProduse() {
    fetch('/api/produse')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('lista-produse');
            tbody.innerHTML = '';
            data.forEach(produs => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${produs.id}</td>
                    <td>${produs.nume}</td>
                    <td>${produs.pret.toFixed(2)} RON</td>
                    <td>${produs.stoc} buc.</td>
                    <td>
                        <button class="btn-edit" data-id="${produs.id}">Editează</button>
                        <button class="btn-delete" data-id="${produs.id}">Șterge</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la preluarea produselor:', error));
}

function handleFormSubmit(event) {
    event.preventDefault();
    const produsData = {
        nume: document.getElementById('nume-produs').value,
        pret: parseFloat(document.getElementById('pret-produs').value),
        stoc: parseInt(document.getElementById('stoc-produs').value, 10)
    };

    let url = '/api/produse/add';
    if (currentEditID !== 0) {
        produsData.id = currentEditID;
        url = '/api/produse/update';
    }

    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(produsData)
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la salvarea produsului'); } return response.json(); })
        .then(data => { console.log(data.mesaj); resetFormular(); loadProduse(); })
        .catch(error => console.error('Eroare formular:', error));
}

function handleEditClick(id) {
    fetch(`/api/produs?id=${id}`)
        .then(response => response.json())
        .then(produs => {
            document.getElementById('nume-produs').value = produs.nume;
            document.getElementById('pret-produs').value = produs.pret;
            document.getElementById('stoc-produs').value = produs.stoc;
            currentEditID = produs.id;
            document.querySelector('#form-add-produs button[type="submit"]').textContent = 'Salvează Modificările';
            window.scrollTo(0, 0);
        })
        .catch(error => console.error('Eroare la preluarea datelor produsului:', error));
}

function handleDelete(id) {
    if (!confirm(`Ești sigur că vrei să ștergi produsul cu ID-ul ${id}?`)) { return; }
    fetch('/api/produse/delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: parseInt(id, 10) })
    })
        .then(response => { if (!response.ok) { throw new Error('Eroare la ștergerea produsului'); } return response.json(); })
        .then(data => { console.log(data.mesaj); loadProduse(); })
        .catch(error => console.error('Eroare la ștergere:', error));
}

function resetFormular() {
    document.getElementById('form-add-produs').reset();
    currentEditID = 0;
    document.querySelector('#form-add-produs button[type="submit"]').textContent = 'Adaugă Produs';
}