window.onload = () => {
    let uuid = parseJwt(window.localStorage.getItem('token')).uuid;
    let username = parseJwt(window.localStorage.getItem('token')).username;

    document.getElementById("uuid").innerText = "(" + uuid + ")";
    document.getElementById("username").innerText = "Username: " + username;
}

const parseJwt = (token) => {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return null;
    }
};