window.onload = function()
{
    rc = Cookie.get('message'); //rc == result cookie

    if (rc.boolean === true) {
        var data = JSON.parse(decodeURIComponent(rc.data));
        alert(data.body.replace('+', ' '));
    }
}