<template>	
	<div class="container-fluid bg-dark py-3 back">
		<div class="jumbotron container">
			<div class="container">
				<h1 class="display-4">{{$lang.messages.hello_world}} </h1>
				<p class="lead">{{$lang.messages.essay}}</p>
				<hr class="my-4">
				<p class="ml-1" >{{$lang.messages.descr}}</p>
				<div class="buttons col-lg-6 col-md-12 col-sm-12 col-xs-12">
					<b-btn variant="info" class="m-1" target="_blank" href="CV_Veronika_Blazhko.pdf"><i class="far fa-file-alt"></i> {{$lang.messages.cv}}</b-btn>
					<b-btn variant="dark" class="m-1 github" target="_blank" href="https://github.com/v-blazhko"><i class="fab fa-github"></i> {{$lang.messages.github}}</b-btn>
					<b-btn variant="light" class="m-1" href="#contact"><i class="far fa-comment"></i> {{$lang.messages.sendmsg}}</b-btn>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: 'aboutme'
}
</script>
<style scoped>
.jumbotron {
	background-color: rgba(57,65,80,0.9);
	color: #fff;
	border-radius: 0px;
	padding: 2rem 2rem;
}

.jumbotron .display-4 {
	padding: 0.5em 0em 1em;
}
.back {
	background-image: url('/img/back.jpg');
	-webkit-background-size: cover;
	-moz-background-size: cover;
	-o-background-size: cover;
	background-size: cover;
}

hr {
	border-color: #ccc;
}


@media (max-width: 768px) {  
	.jumbotron {
		padding: 0em 0em 2em 0em;
	}

	.jumbotron .display-4 {
		padding: 0.5em 0em 0.1em;
	}



}

.buttons {
	display: flex;
	flex-wrap: wrap;
	padding: 0px;
	margin: 0px;
}

.btn {
	/*border: 1px solid #888; */
	flex-shrink: 1;
	flex-grow: 1;
}

.btn-dark:hover {
    color: #fff;
    background-color: #202428;
    border-color: #202428;
}

.btn-dark {
    color: #fff;
    background-color: #292E33;
    border-color: #292E33;
}



</style>