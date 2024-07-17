## Snippetbox

This is the final code of the [Let's Go e-book by Alex Edwards](https://lets-go.alexedwards.net/). It's a full-stack Go web application called "Snippetbox" that lets users CRUD text snippets (similar to GitHub gists).

<img width="500" src="./lets-go-screenshot.png" />

### Features
- RESTful routing (Go 1.22’s HTTP Package)
- HTML Templating
- Authentication and Authorization
- Level logging and centralized error handling
- Middlewares
- Session Management
- MySQL database
- Security (HTTPS,OWASP Secure Heards and CSRF)
- Testing

### API
<table>
<thead>
<tr>
<th>Method</th>
<th>Pattern</th>
<th>Action</th>
</tr>
</thead>

<tbody>
<tr>
<td>GET</td>
<td>/</td>
<td>Display the home page</td>
</tr>

<tr>
<td>GET</td>
<td><span>/snippet/view/{id}</span></td>
<td>Display a specific snippet</td>
</tr>

<tr>
<td>GET</td>
<td>/snippet/create</td>
<td>Display a HTML form for creating a new snippet</td>
</tr>

<tr>
<td>POST</td>
<td>/snippet/create</td>
<td>Create a new snippet</td>
</tr>

<tr>
<td>GET</td>
<td>/user/signup</td>
<td>Display a HTML form for signing up a new user</td>
</tr>

<tr>
<td>POST</td>
<td>/user/signup</td>
<td>Create a new user</td>
</tr>

<tr>
<td>GET</td>
<td>/user/login</td>
<td>Display a HTML form for logging in a user</td>
</tr>

<tr>
<td>POST</td>
<td>/user/login</td>
<td>Authenticate and login the user</td>
</tr>

<tr>
<td>POST</td>
<td>/user/logout</td>
<td>Logout the user</td>
</tr>

<tr>
<td>GET</td>
<td>/account/view</td>
<td>View account details</td>
</tr>

<tr>
<td>GET</td>
<td>/account/password/update</td>
<td>Display a HTML form for updating password</td>
</tr>

<tr>
<td>POST</td>
<td>/account/password/update</td>
<td>Update the user password</td>
</tr>

<tr>
<td>GET</td>
<td>/about</td>
<td>About page</td>
</tr>

</tbody>
</table>

