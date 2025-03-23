(() => {
    const year = document.querySelector('#year');
    if (year) year.innerText = ((new Date()).getFullYear()).toString();
})();