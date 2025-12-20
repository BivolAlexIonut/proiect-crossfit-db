document.addEventListener("DOMContentLoaded", () => {
    loadMentorat();
    loadAntrenoriSelect();
    loadMembriSelect();

    document.getElementById("form-add-mentorat").addEventListener("submit", async (e) => {
        e.preventDefault();
        const antrenorID = document.getElementById("select-antrenor").value;
        const membruID = document.getElementById("select-membru").value;

        try {
            const res = await fetch("/api/mentorat/add", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    antrenorID: parseInt(antrenorID),
                    membruID: parseInt(membruID)
                })
            });
            if (!res.ok) throw new Error("Eroare la adăugare");
            alert("Relație de mentorat creată!");
            loadMentorat();
        } catch (err) {
            console.error(err);
            alert("Nu s-a putut crea relația (posibil duplicat).");
        }
    });
});

async function loadMentorat() {
    const tbody = document.getElementById("lista-mentorat");
    tbody.innerHTML = "<tr><td colspan='3'>Se încarcă...</td></tr>";
    try {
        const res = await fetch("/api/mentorat");
        const data = await res.json();
        tbody.innerHTML = "";
        if (!data || data.length === 0) {
            tbody.innerHTML = "<tr><td colspan='3'>Nu există relații de mentorat.</td></tr>";
            return;
        }
        data.forEach(m => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${m.numeAntrenor}</td>
                <td>${m.numeMembru}</td>
                <td>
                    <button onclick="deleteMentorat(${m.antrenorID}, ${m.membruID})">Șterge</button>
                </td>
            `;
            tbody.appendChild(row);
        });
    } catch (err) {
        console.error(err);
        tbody.innerHTML = "<tr><td colspan='3'>Eroare la încărcare.</td></tr>";
    }
}

async function deleteMentorat(antrenorID, membruID) {
    if (!confirm("Sigur ștergi această relație?")) return;
    try {
        const res = await fetch("/api/mentorat/delete", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ antrenorID: antrenorID, membruID: membruID })
        });
        if (!res.ok) throw new Error("Eroare la ștergere");
        loadMentorat();
    } catch (err) {
        console.error(err);
        alert("Eroare la ștergere.");
    }
}

async function loadAntrenoriSelect() {
    const select = document.getElementById("select-antrenor");
    try {
        const res = await fetch("/api/antrenori");
        const data = await res.json();
        data.forEach(a => {
            const opt = document.createElement("option");
            opt.value = a.id;
            opt.textContent = `${a.nume} ${a.prenume}`;
            select.appendChild(opt);
        });
    } catch (err) { console.error(err); }
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
