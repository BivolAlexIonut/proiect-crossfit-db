/**
 * =============================================================================
 * SCRIPT MANAGEMENT ACHIZIȚII
 * =============================================================================
 */

let achizitiiList = [];
let sortDirection = 1;
let lastSortColumn = '';

document.addEventListener("DOMContentLoaded", () => {
    loadAchizitii();
    loadMembriSelect();
    loadProduseSelect();

    document.getElementById("form-add-achizitie").addEventListener("submit", handleAddAchizitie);
});

// --- LOAD DATA ---

async function loadAchizitii() {
    try {
        const res = await fetch("/api/achizitii");
        if (!res.ok) throw new Error("Eroare rețea");
        achizitiiList = await res.json() || [];
        renderAchizitii();
    } catch (err) {
        alert("Eroare la încărcarea achizițiilor: " + err.message);
    }
}

async function loadMembriSelect() {
    try {
        const res = await fetch("/api/membri");
        const data = await res.json() || [];
        const select = document.getElementById("select-membru");
        select.innerHTML = '<option value="">Alege membrul...</option>';
        
        data.forEach(m => {
            const opt = document.createElement("option");
            opt.value = m.id;
            opt.textContent = `${m.nume} ${m.prenume}`;
            select.appendChild(opt);
        });
    } catch (err) {
        console.error("Nu s-au putut încărca membrii", err);
    }
}

async function loadProduseSelect() {
    try {
        const res = await fetch("/api/produse");
        const data = await res.json() || [];
        const select = document.getElementById("select-produs");
        select.innerHTML = '<option value="">Alege produsul...</option>';
        
        data.forEach(p => {
            const opt = document.createElement("option");
            opt.value = p.id;
            opt.textContent = `${p.nume} (Stoc: ${p.stoc}, Preț: ${p.pret} RON)`;
            // Dezactivăm opțiunea dacă nu e stoc, vizual
            if (p.stoc <= 0) {
                opt.textContent += " - INDISPONIBIL";
                opt.disabled = true;
            }
            select.appendChild(opt);
        });
    } catch (err) {
        console.error("Nu s-au putut încărca produsele", err);
    }
}

// --- ACTIONS ---

async function handleAddAchizitie(e) {
    e.preventDefault();
    const membruID = document.getElementById("select-membru").value;
    const produsID = document.getElementById("select-produs").value;
    const cantitateVal = document.getElementById("cantitate").value;
    const cantitate = parseInt(cantitateVal, 10);

    if (isNaN(cantitate) || cantitate <= 0) {
        alert("Eroare: Cantitatea trebuie să fie strict mai mare decât 0.");
        return;
    }

    try {
        const res = await fetch("/api/achizitii/add", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                membruID: parseInt(membruID),
                produsID: parseInt(produsID),
                cantitate: cantitate
            })
        });

        if (!res.ok) {
            const errText = await res.text();
            throw new Error(errText);
        }

        alert("Achiziție înregistrată!");
        // Reîncărcăm tot pentru a reflecta stocurile noi
        loadAchizitii();
        loadProduseSelect();
        e.target.reset(); // Reset form
    } catch (err) {
        alert("Eroare: " + err.message);
    }
}

async function deleteAchizitie(id) {
    if (!confirm("Sigur ștergi această achiziție? Stocul va fi restituit.")) return;

    try {
        const res = await fetch("/api/achizitii/delete", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id: id })
        });

        if (!res.ok) {
            const errText = await res.text();
            throw new Error(errText);
        }

        alert("Achiziție ștearsă.");
        loadAchizitii();
        loadProduseSelect();
    } catch (err) {
        alert("Eroare la ștergere: " + err.message);
    }
}

// --- RENDER & SORT ---

function renderAchizitii() {
    const tbody = document.getElementById("lista-achizitii");
    tbody.innerHTML = "";
    if (achizitiiList.length === 0) {
        tbody.innerHTML = "<tr><td colspan='7'>Nu există achiziții.</td></tr>";
        return;
    }
    achizitiiList.forEach(a => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td data-label="ID">${a.id}</td>
            <td data-label="Membru">${a.numeMembru}</td>
            <td data-label="Produs">${a.numeProdus}</td>
            <td data-label="Data">${a.dataAchizitiei}</td>
            <td data-label="Cantitate">${a.cantitate}</td>
            <td data-label="Total">${a.pretTotal ? a.pretTotal.toFixed(2) : "0.00"}</td>
            <td data-label="Acțiuni">
                <button class="btn-delete" onclick="deleteAchizitie(${a.id})">Șterge</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

function sortAchizitii(column) {
    if (lastSortColumn === column) {
        sortDirection *= -1;
    } else {
        sortDirection = 1;
        lastSortColumn = column;
    }

    achizitiiList.sort((a, b) => {
        let valA = a[column];
        let valB = b[column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return -1 * sortDirection;
        if (valA > valB) return 1 * sortDirection;
        return 0;
    });

    renderAchizitii();
}