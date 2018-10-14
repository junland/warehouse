package server

// deftmpl defines the default html template
const deftmpl = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
  <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
</head>
	<body>
	 <div class="container">
	 <table class="centered">
	  <thead>
	    <tr>
		   <th>Name</th>
		   <th>Last modified</th>
		   <th>Size</th>
	    </tr>
	  </thead>
	<tbody>
	 <tr>
	    <td><a href="../">Go Back...</a></td>
	 </tr>
	{{range .Items}}
	  <tr>
	    <td><a href="{{.Name}}">{{.Name}}</a></td>
	    <td><a>{{.LastMod}}</a></td>
	    <td><a>{{.Size}}</a></td>
	  </tr>
	{{end}}
	</tbody>
	 </table>
	 </div>
    </body>
</html>`
