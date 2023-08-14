window.onload = function() {
    console.log("loaded nofriends.js");

    getActivePublicServers().then((servers) => {
        createTable(servers);
    });

}


async function getActivePublicServers() {
    const response = await fetch('/api/v1/server/getActivePublicServers');
    const data = await response.json();
    if (data.success == true) {
        console.log("getActivePublicServers() success");
        return data.servers;
    }
    return {};  // return something else later?
}

function createTable(servers) {
    let table = document.getElementById("public-servers-table");
    let tbody = table.createTBody();

    // TODO all this can probably be moved to html file
    let thead = table.createTHead();
    let headerRow = thead.insertRow();
    let mapHeader = headerRow.insertCell(0);
    let playersHeader = headerRow.insertCell(1);
    let urlHeader = headerRow.insertCell(2);
    mapHeader.innerHTML = "Map";
    playersHeader.innerHTML = "Players";
    urlHeader.innerHTML = "Connect";

    if (servers.length == 0) {
        let row = tbody.insertRow();
        let noServers = row.insertCell(0);
        noServers.innerHTML = "No public servers! Sorry friend.";
        return;
    }

    servers.forEach(server => {
        let row = tbody.insertRow();
        let map = row.insertCell(0);
        let players = row.insertCell(1);
        let url = row.insertCell(2);
        
        map.innerHTML = server.starting_map;
        players.innerHTML = server.players + "/" + server.maxplayers;
        url.innerHTML = server.url;
        // join.innerHTML = "<button type=\"button\" onclick=\"joinServer('" + server.ip + "', '" + server.port + "')\">Join</button>";
    });
}
