document.addEventListener("DOMContentLoaded", () => {
    loadCompetitii();
    loadMembriSelect();
    loadCompetitiiSelect();
    loadParticipari();

    // Formular Adaugare Competitie
    document.getElementById("form-add-competitie").addEventListener("submit", async (e) => {
        e.preventDefault();
        const data = {
            nume: document.getElementById("nume").value,
            data: document.getElementById("data").value,
            locatie: document.getElementById("locatie").value,
            taxa: parseFloat(document.getElementById("taxa").value)
        };

        try {
            const res = await fetch("/api/competitii/add", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data)
            });
            if (!res.ok) throw new Error("Eroare adaugare competitie");
            alert("Competiție creată!");
            loadCompetitii();
            loadCompetitiiSelect();
        } catch (err) {
            console.error(err);
            alert("Eroare la creare.");
        }
    });

    // Formular Inscriere Membru la Competitie
    document.getElementById("form-inscriere-comp").addEventListener("submit", async (e) => {
        e.preventDefault();
        const data = {
            competitieID: parseInt(document.getElementById("select-competitie").value),
            membruID: parseInt(document.getElementById("select-membru").value)
        };

        try {
            const res = await fetch("/api/competitii/participari/add", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data)
            });
            if (!res.ok) throw new Error("Eroare inscriere");
            alert("Membru înscris cu succes!");
            loadParticipari();
        } catch (err) {
            console.error(err);
            alert("Eroare la înscriere (poate e deja înscris).");
        }
    });
});

async function loadCompetitii() {
    const tbody = document.getElementById("lista-competitii");
    try {
        const res = await fetch("/api/competitii");
        const data = await res.json();
        tbody.innerHTML = "";
        if (!data) return;
        data.forEach(c => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${c.id}</td>
                <td>${c.nume}</td>
                <td>${c.data}</td>
                <td>${c.locatie}</td>
                <td>${c.taxa} RON</td>
                <td>
                    <button onclick="deleteCompetitie(${c.id})">Șterge</button>
                </td>
            `;
            tbody.appendChild(row);
        });
    } catch (err) { console.error(err); }
}

async function loadParticipari() {
    const tbody = document.getElementById("lista-participari");
    try {
        const res = await fetch("/api/competitii/participari");
        const data = await res.json();
        tbody.innerHTML = "";
        if (!data) return;
        data.forEach(p => {
            const row = document.createElement("tr");
            const loc = p.loculObtinut > 0 ? p.loculObtinut : "-";
            row.innerHTML = `
                <td>${p.numeCompetitie}</td>
                <td>${p.numeMembru}</td>
                <td>${loc}</td>
            `;
            tbody.appendChild(row);
        });
    } catch (err) { console.error(err); }
}

async function deleteCompetitie(id) {
    if (!confirm("Sigur ștergi competiția? Se vor șterge și toate înscrierile!")) return;
    try {
        const res = await fetch("/api/competitii/delete", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id: id })
        });
        if (res.ok) {
            loadCompetitii();
            loadCompetitiiSelect();
            loadParticipari();
        }
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

async function loadCompetitiiSelect() {
    const select = document.getElementById("select-competitie");
    select.innerHTML = '<option value="">Alege competiția...</option>';
    try {
        const res = await fetch("/api/competitii");
        const data = await res.json();
        if(!data) return;
        data.forEach(c => {
            const opt = document.createElement("option");
            opt.value = c.id;
            opt.textContent = `${c.nume} (${c.data})`;
            select.appendChild(opt);
        });
    } catch (err) { console.error(err); }
}
