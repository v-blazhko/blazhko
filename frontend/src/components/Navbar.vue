<template>

	<nav class="navbar navbar-expand-lg navbar-light bg-light sticky-top">
		<div class="container">
			<a class="navbar-brand" href="#"><img class="img-fluid logo" src="img/black.svg">  blazhko.tech</a>
			<button class="navbar-toggler" type="button" data-toggle="collapse" v-on:click="toggleNav()" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>

			<div class="collapse navbar-collapse" id="navbarSupportedContent">
				<ul class="navbar-nav ml-auto">
					<li class="nav-item m-0">
						<a class="nav-link" v-on:click="$emit('en')">EN</a>
					</li>
					<li class="nav-item m-0">
						<a class="nav-link" v-on:click="$emit('ru')">RU</a>
					</li>
				</ul>
			</div>
		</div>
	</nav>
</template>

<script>

export default {
	name: 'navbar',
	methods: {
		toggleNav: function () {
			var x = document.getElementById("navbarSupportedContent");
			if (x.style.display === "none") {
				x.style.display = "block";
			} else {
				x.style.display = "none";
			}
		}
	}

}
</script>

<style scoped>

.logo {
	height: 2.2em;
}

.nav-link {
	cursor: pointer;
}

.navbar-brand {
	font-family: 'Courier new';
	font-size: 1.15em;
	font-weight: bold;
}
</style>