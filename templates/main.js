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

const allowedCombinations = {
    "audio-source-int-ext": "audio-dest-int-ext",
    "audio-source-ext-main": "audio-dest-ext-main",
    "button-meta-source": "button-meta-dest"
}

let selectedSourceText = "";
let selectedSourceID = "";

let connections = {};
let undoConnections = {};

let sourceButtons = Array.from(document.getElementsByClassName("button-source"));
let audioSourceButtons = Array.from(document.getElementsByClassName("button-audio-source"));
let metaSourceButtons = Array.from(document.getElementsByClassName("button-meta-source"));
let destinationButtons = Array.from(document.getElementsByClassName("button-dest"));

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

        connections[e.innerText] = {
            type: selectedSourceType,
            shortName: selectedSource.innerText,
            data: selectedSourceType == "META" ? metadataGroups[selectedMetaGroup].Members[getMetaBtnId(selectedSource) - 1].LongName : ""
        };

        document.getElementById("connections").innerHTML = "";
        for (c in connections) {
            let row = document.createElement("li");
            row.innerText = `${connections[c].shortName} -> ${c}`;
            document.getElementById("connections").appendChild(row);
        }
    }
})

document.getElementById("take-button").onclick = () => {
    fetch("/control/take", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(connections)
    }).then((r) => r.json()).then((d) => {
        for (destination in d) {
            document.getElementById(`dest-connection-${destination.replace(" ", "-")}`).innerText = d[destination];
        }
    });

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

let paginationData = {};
fetch("/control/meta").then((r) => r.json()).then((d) => {
    paginationData = d;
})

{{ range .Sources }}
    {{ range . }}
        {{ if and (.ID) (not (eq .ButtonType "POSTPAGE")) }}
            {{ if eq .ButtonType "RAW"}}
                document.getElementById("{{.ID}}").onclick = () => {
                    selectedSourceText = "{{.Name}}";
                    selectedSourceID = "{{.ID}}";
                    sourceButtons.forEach((b) => {
                        b.style.border = "";
                    })
                    document.getElementById("{{.ID}}").style.border = "5px solid red";
                }
            {{ else if eq .ButtonType "PAGE"}}
                document.getElementById("{{.ID}}").onclick = () => {
                    let pageData = paginationData.Sources["{{.PageGroup}}"]["{{.Name}}"];

                    for (let i = 0; i < pageData.length; i++) {
                        let button = document.getElementById(`{{.PageGroup}}-${i+1}`);
                        button.innerText = pageData[i].Name;
                        button.style.visibility = pageData.length > [i] == "" ? "hidden": "visible";
                        button.onclick = () => {
                            selectedSourceText = pageData[i].Name;
                            selectedSourceID = pageData[i].ID;
                            sourceButtons.forEach((b) => {
                                b.style.border = "";
                            })
                            button.style.border = "5px solid red";      
                        };
                    }
                }
            {{ end }}

        {{ end }}
    {{ end }}
{{ end }}

fetch("/control/user").then((r) => r.text()).then((d) => {
    document.getElementById("user").innerText = `MCR Operator: ${d}`;
})