package server

// deftmpl defines the default html template
const deftmpl = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
	<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.4.1/css/all.css" integrity="sha384-5sAR7xN1Nv6T6+dT2mhtzEpVJvfS3NScPQTrOxhwjIuvcA67KV2R5Jz6kr4abQsz" crossorigin="anonymous">
	<link rel="shortcut icon" href="data:image/x-icon;," type="image/x-icon"> 
</head>
	<body>
	 <div class="container">
	 <div class="col-md-8 col-xs-12 vbottom">
			<h4>Index of {{.RelPath}}</h4>
	 </div>
	 <table class="responsive-table">
	  <thead>
	    <tr>
		   <th>Name</th>
		   <th>Last modified</th>
		   <th>Size</th>
	    </tr>
	  </thead>
	<tbody>
	 <tr>
	    <td><i class="fas fa-reply"></i><a href="../">  Go Back...</a></td>
	 </tr>
	{{range .Items}}
		<tr>
			{{if .Dir}}
       <td><i class="fas fa-folder-open"></i><a href="{{.Name}}"> {{.Name}}</a></td>
			{{else}}
			 <td><i class="fas fa-file"></i><a href="{{.Name}}"> {{.Name}}</a></td>
			{{end}}
	    <td><a>{{.LastMod}}</a></td>
	    <td><a>{{.HumanSize}}</a></td>
	  </tr>
	{{end}}
	</tbody>
	 </table>
	 </div>
    </body>
</html>`
