<template>	
	<div class="container-fluid bg-dark py-3 back">
		<div class="jumbotron container">
			<div class="container">
				<h1 class="display-4">{{$lang.messages.hello_world}} </h1>
				<p class="lead">{{$lang.messages.essay}}</p>
				<hr class="my-4">
				<p>{{$lang.messages.descr}}</p>
				<b-btn variant="info" class="m-1" target="_blank" href="static/CV_Veronika_Blazhko.pdf"><i class="far fa-file-alt"></i> {{$lang.messages.cv}}</b-btn>
				<b-btn variant="dark" class="m-1" target="_blank" href="https://github.com/v-blazhko"><i class="fab fa-github"></i> {{$lang.messages.github}}</b-btn>
				<b-btn variant="light" class="m-1" href="#contact"><i class="far fa-comment"></i> {{$lang.messages.sendmsg}}</b-btn>
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
	background-color: rgba(57,65,80,0.8);
	color: #fff;
	border-radius: 0px;
	padding: 2rem 2rem;
}

.jumbotron .display-4 {
	padding: 0.5em 0em 1em;
}
.back {
	background-image: url('/static/img/back.jpg');
	-webkit-background-size: cover;
	-moz-background-size: cover;
	-o-background-size: cover;
	background-size: cover;
}

hr {
	border-color: #ccc;
}

</style>
