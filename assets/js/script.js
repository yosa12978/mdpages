function renderMarkdown(id, md) {
    let purified = DOMPurify.sanitize(marked.parse(md), { ADD_TAGS: ["iframe"], ADD_ATTR: ['allow', 'allowfullscreen', 'frameborder', 'scrolling'] });
    document.getElementById(id).innerHTML = purified;
}

function getDate(dateISO8601, timezone = 'GMT') {
    const d = new Date(dateISO8601);
    return d.getFullYear().toString() + "/" + (d.getMonth() + 1).toString() + "/" + d.getDate().toString();
}