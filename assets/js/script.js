function renderMarkdown(id, md) {
    document.getElementById(id).innerHTML = DOMPurify.sanitize(marked.parse(md));
}