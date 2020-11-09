<template>
	<div>
		
		<vue-star animate="animated bounceIn" color="#ffff4d">
			<a slot="icon" class="fa fa-star game"></a>
		</vue-star>
		<h1 class="text-center mt-5">If you've scrolled down till this point...</h1>
		<p class="text-center">..there is a game for you!</p>
		<section id="application">

			<div id="div1">
				<button type="submit" @click="updateCount()" id="clickButton">
				</button>
				<img id="img1" src="static/img/giphy.gif" >
			</div>

			<!-- Second status  -->
			<div id="status"></div>
			<h2> You clicked <b>{{ numClicks }}</b> times</h2>
			<h2> You have <b>{{ secs }}</b> secs left </h2>

		</section>
	</div>
</template>

<style >
.game {
	color: #5d758c;
	font-size: 3em;
}
</style>
<script>
	import VueStar from 'vue-star'
	export default {

		name: 'game',

		data() { 
			return {
				
				numClicks: 0,
				secs: 10,
				colors:['#9bcd77','#d777b0','#fed401','#ff2c3c', '#66c652'],
				width: 200
			}

		},

		components: {
			VueStar
		},

		mounted: function(){
			var self = this;
			setInterval(function(){
				if(self.active){
					self.secs--;
					if(self.secs == 0){
						self.active = false;
						var button = document.getElementById("div1");
						button.classList.add("over");
						this.width = 200;
						button.style.width = this.width + 'px';
						button.style.height = this.width + 'px';
						var img = document.getElementById("img1");
						img.style.display = "block";
						var stat = document.getElementById("status");
						stat.innerHTML = '<h2>Game Over</h2>';
					} else {
						var stat = document.getElementById("status");
						stat.innerHTML = '<h2></h2>';
					}
				}
			},1000);
		},

		methods: {
			updateCount: function () {
				this.numClicks += 1;    
				if(this.active == false){
					var button = document.getElementById("clickButton");
					var stat = document.getElementById("status");
					stat.innerHTML = '';
					button.innerHTML = '';
					this.secs = 10;
					this.numClicks = 0;
					this.active = true;
				}

				var color = this.colors[ Math.round( Math.random()*(this.colors.length-1) ) ];
				console.log(color);
				var button = document.getElementById("clickButton");
				button.style.backgroundColor = color;
				console.log(button.style.width);
				this.width = this.width + 1;
				button.style.width = this.width + 'px';
				button.style.height = this.width + 'px';
			}
		}
	}
</script>

<style scoped>
body {
	background-color: #323133;
	color: #fff;
	text-align: center;
	font-family: 'Dosis', sans-serif; background-image:url('http://imgh.us/bg_12.svg');
	background-repeat: no-repeat;
	background-size: cover;
	height: 100vh;
	background-attachment: fixed;
	min-height: 600px;
}

button {

	border: none;
	color: #fff;
	width: 200px;
	height: 200px;
	transition: 0.3s;
	border-radius: 50%;
	background-color: #f2ca27;
}

button:active {
	width: 200px;
	height: 200px;
}
button:focus {
	outline: none;
}
#status {
	background: transparent;
	
}
#status h2 {
	font-size: 60px;
	text-transform: uppercase;
}
#timer {
	text-align: center;
	margin-top: 140px;
}

#application {
	margin: 0 auto;
}
#application h1 {
	color: #fff;
	margin-top: 140px;
	font-weight: 300;
}
.over {
	background-color: transparent;
	
	height: 200px;
	width: 200px;
}

#img1 {
	display: none;
}
</style>
