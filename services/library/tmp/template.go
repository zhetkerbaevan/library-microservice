package tmp

var BooksTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Library</title>
</head>
<body>
    <h1>Books List</h1>
    <table border="1">
        <tr>
            <th>Book ID</th>
            <th>Name</th>
            <th>Author</th>
            <th>Genre</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Id}}</td>
            <td>{{.Name}}</td>
            <td>{{.Author}}</td>
            <td>{{.Genre}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
