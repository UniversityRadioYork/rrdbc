setInterval(() => {
    let d = new Date();
    document.getElementById("clock").innerText = `${String(
        d.getHours()
    ).padStart(2, "0")}:${String(d.getMinutes()).padStart(2, "0")}:${String(
        d.getSeconds()
    ).padStart(2, "0")}`;
}, 1000);