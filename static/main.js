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