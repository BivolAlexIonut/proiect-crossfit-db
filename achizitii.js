document.addEventListener("DOMContentLoaded", () => {
    loadAchizitii();
    loadMembriSelect();
    loadProduseSelect();

    document.getElementById("form-add-achizitie").addEventListener("submit", async (e) => {
        e.preventDefault();
        const membruID = document.getElementById("select-membru").value;
        const produsID = document.getElementById("select-produs").value;
        const cantitate = document.getElementById("cantitate").value;

        try {
            const res = await fetch("/api/achizitii/add", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    membruID: parseInt(membruID),
                    produsID: parseInt(produsID),
                    cantitate: parseInt(cantitate)
                })
            });
            if (!res.ok) throw new Error("Eroare la adăugare");
            alert("Achiziție înregistrată!");
            loadAchizitii();
        } catch (err) {
            console.error(err);
            alert("Nu s-a putut adăuga achiziția.");
        }
    });
});

async function loadAchizitii() {
    const tbody = document.getElementById("lista-achizitii");
    tbody.innerHTML = "<tr><td colspan='6'>Se încarcă...</td></tr>";
    try {
        const res = await fetch("/api/achizitii");
        const data = await res.json();
        tbody.innerHTML = "";
        if (!data || data.length === 0) {
            tbody.innerHTML = "<tr><td colspan='6'>Nu există achiziții.</td></tr>";
            return;
        }
        data.forEach(a => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${a.id}</td>
                <td>${a.numeMembru}</td>
                <td>${a.numeProdus}</td>
                <td>${a.dataAchizitiei}</td>
                <td>${a.cantitate}</td>
                <td>
                    <button onclick="deleteAchizitie(${a.id})">Șterge</button>
                </td>
            `;
            tbody.appendChild(row);
        });
    } catch (err) {
        console.error(err);
        tbody.innerHTML = "<tr><td colspan='6'>Eroare la încărcare.</td></tr>";
    }
}

async function deleteAchizitie(id) {
    if (!confirm("Sigur ștergi această achiziție?")) return;
    try {
        const res = await fetch("/api/achizitii/delete", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id: id })
        });
        if (!res.ok) throw new Error("Eroare la ștergere");
        loadAchizitii();
    } catch (err) {
        console.error(err);
        alert("Eroare la ștergere.");
    }
}

async function loadMembriSelect() {
    const select = document.getElementById("select-membru");
    try {
        const res = await fetch("/api/membri");
        const data = await res.json();
        data.forEach(m => {
            const opt = document.createElement("option");
            opt.value = m.id;
            opt.textContent = `${m.nume} ${m.prenume}`;
            select.appendChild(opt);
        });
    } catch (err) { console.error(err); }
}

async function loadProduseSelect() {
    const select = document.getElementById("select-produs");
    try {
        const res = await fetch("/api/produse");
        const data = await res.json();
        data.forEach(p => {
            const opt = document.createElement("option");
            opt.value = p.id;
            opt.textContent = `${p.nume} (Stoc: ${p.stoc})`;
            select.appendChild(opt);
        });
    } catch (err) { console.error(err); }
}
