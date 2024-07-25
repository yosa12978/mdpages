function renderMarkdown(id, md) {
    document.getElementById(id).innerHTML = DOMPurify.sanitize(marked.parse(md));
}

function getDate(dateISO8601, timezone = 'GMT') {
    const d = new Date(dateISO8601);
    return d.getFullYear().toString() + "/" + (d.getMonth() + 1).toString() + "/" + d.getDate().toString();
}