var groupBy = function (xs, key) {
    return xs.reduce(function (rv, x) {
        (rv[x[key]] = rv[x[key]] || []).push(x);
        return rv;
    }, {});
};
function convert(unixtimestamp) {

    // Months array
    var months_arr = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

    // Convert timestamp to milliseconds
    var date = new Date(unixtimestamp);

    // Year
    var year = date.getFullYear();

    // Month
    var month = months_arr[date.getMonth()];

    // Day
    var day = date.getDate();

    // Hours
    var hours = date.getHours();

    // Minutes
    var minutes = "0" + date.getMinutes();

    // Seconds
    var seconds = "0" + date.getSeconds();

    // Display date time in MM-dd-yyyy h:m:s format
    var convdataTime = month + '-' + day + '-' + year + ' ' + hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2);

    return convdataTime;

}

function cleanUpAddress(address) {
    var fixAddress
    fixAddress = address.replace('http://', '');
    fixAddress = fixAddress.replace('/_read', '');
    return fixAddress
}
function isEmptyOrSpaces(str) {
    return str === null || str.match(/^ *$/) !== null;
}