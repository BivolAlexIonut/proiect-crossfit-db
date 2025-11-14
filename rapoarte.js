window.addEventListener('DOMContentLoaded', () => {
    loadRaportAbonamente();
    loadRaportVizualizare();
    loadRaportComplex();
});

// Raport 1: Popularitate Abonamente (GROUP BY / HAVING)
function loadRaportAbonamente() {
    fetch('/api/raport/abonamente')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('raport-abonamente');
            tbody.innerHTML = '';
            if (!data) data = [];
            data.forEach(item => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${item.tipAbonament}</td>
                    <td>${item.numarMembri}</td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la raportul abonamente:', error));
}

// Raport 2: Vizualizare Membri-Abonamente (VIEW)
function loadRaportVizualizare() {
    fetch('/api/raport/vizualizare-membri')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('raport-vizualizare');
            tbody.innerHTML = '';
            if (!data) data = [];
            data.forEach(item => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${item.nume}</td>
                    <td>${item.prenume}</td>
                    <td>${item.email}</td>
                    <td>${item.tipAbonament}</td>
                    <td>${item.pret.toFixed(2)} RON</td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la raportul vizualizare:', error));
}

// Raport 3: Interogare Complexă (JOIN pe 5 tabele)
function loadRaportComplex() {
    fetch('/api/raport/complex-inscrieri')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('raport-complex');
            tbody.innerHTML = '';
            if (!data) data = [];
            data.forEach(item => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${item.numeMembru}</td>
                    <td>${item.prenumeMembru}</td>
                    <td>${item.numeClasa}</td>
                    <td>${item.numeAntrenor}</td>
                `;
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Eroare la raportul complex:', error));
}