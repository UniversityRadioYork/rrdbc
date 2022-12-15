setInterval(() => {
    let d = new Date();
    document.getElementById("clock").innerText = `${String(
        d.getHours()
    ).padStart(2, "0")}:${String(d.getMinutes()).padStart(2, "0")}:${String(
        d.getSeconds()
    ).padStart(2, "0")}`;
}, 1000);

const setStatus = (text) => {
    document.getElementById("status").innerText = text;
    setTimeout(() => {
        document.getElementById("status").innerText = "";
    }, 5000)
}

let selectedSource = null;
let sourceButtons = Array.from(document.getElementsByClassName("button-source"));

sourceButtons.forEach((e) => {
    e.onclick = () => {
        selectedSource = e;
        sourceButtons.forEach((b) => {
            b.style.border = "";
        })
        e.style.border = "5px solid red";
    }
})

let destinationButtons = Array.from(document.getElementsByClassName("button-dest"));
const allowedCombinations = {
    "audio-source-int-ext": "audio-dest-int-ext",
    "audio-source-ext-main": "audio-dest-ext-main",
    "button-meta-source": "button-meta-dest"
}
let connections = {};
let undoConnections = {};

destinationButtons.forEach((e) => {
    e.onclick = () => {
        if (selectedSource == null) {
            setStatus("You must select a source.")
            return;
        }

        for (k in allowedCombinations) {
            if (selectedSource.classList.contains(k) && !e.classList.contains(allowedCombinations[k])) {
                setStatus(`You can't connect ${selectedSource.innerText} to ${e.innerText}`);
                return;
            }
        }

        connections[e.innerText] = selectedSource.innerText;

        document.getElementById("connections").innerHTML = "";
        for (c in connections) {
            let row = document.createElement("li");
            row.innerText = `${connections[c]} -> ${c}`;
            document.getElementById("connections").appendChild(row);
        }
    }
})

document.getElementById("take-button").onclick = () => {
    // TODO Send the Request

    document.getElementById("undo-connections").innerHTML = document.getElementById("connections").innerHTML;
    document.getElementById("connections").innerHTML = "";
    sourceButtons.forEach((b) => {
        b.style.border = "";
    })

    undoConnections = connections;
    connections = {};
    selectedSource = null;
}

document.getElementById("undo-button").onclick = () => {
    // TODO Send the request

    document.getElementById("undo-connections").innerHTML = "";
    undoConnections = {};
}

let metadataGroups = [];

fetch("/control/meta").then((r) => r.json()).then((d) => {
    metadataGroups = d;
    for (let i = 0; i < 16; i++) {
        let grpBtn = document.getElementById(`metagrp-${i + 1}`);
        grpBtn.innerText = metadataGroups[i].GrpName;

        if (metadataGroups[i].GrpName == "") {
            grpBtn.style.visibility = "hidden";
            continue;
        }

        grpBtn.onclick = () => {
            for (let j = 0; j < 16; j++) {
                document.getElementById(`metabtn-${j + 1}`).innerText = metadataGroups[i].Members[j].ShortName;
                document.getElementById(`metabtn-${j + 1}`).style.visibility = metadataGroups[i].Members[j].ShortName == "" ? "hidden" : "visible";
            }
            Array.from(document.getElementsByClassName("button-meta-source")).forEach((b) => {
                b.style.border = "";
            })
            Array.from(document.getElementsByClassName("button-meta-group")).forEach((b) => {
                b.style.border = "";
            })
            grpBtn.style.border = "5px solid red";
        }

    }
})

fetch("/control/user").then((r) => r.text()).then((d) => {
    document.getElementById("user").innerText = `MCR User: ${d}`;
})