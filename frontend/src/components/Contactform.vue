<template>
	<div id="contact">
		<h1 class="text-center">{{$lang.messages.sendmsg}}</h1>
		<div class="row">
			<div class="col-sm-12 col-lg-4">
				<ul class="row" style="display:flex;">
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fas fa-map-marker-alt mr-2"></i>   Location: <a href="https://www.google.fi/maps/place/%D0%A2%D0%B0%D0%BC%D0%BF%D0%B5%D1%80%D0%B5/@61.6319675,23.5501175,10z/data=!4m5!3m4!1s0x468edf554593da5d:0x6adfe3bd1e0b22c0!8m2!3d61.4977524!4d23.7609535" target="_blank">Tampere</a></li>
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fa fa-envelope mr-2"></i>   Email: <a href="mailto:veronica@blazhko.tech" target="_blank">veronica@blazhko.tech</a></li>
					<li  class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-linkedin-in mr-2"></i>   LinkedIn: <a href="https://linkedin.com/in/v-blazhko/" target="_blank">v-blazhko</a></li>
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-github mr-2"></i>   Github: <a href="https://github.com/v-blazhko/" target="_blank">v-blazhko</a></li>
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-vk mr-2"></i>   VK: <a href="https://vk.com/v_blazhko" target="_blank">v_blazhko</a></li>
				</ul>
			</div>
			<div class="col-sm-12 col-lg-8">
				<div class="row">
					<div class="col-md-12 mx-auto">
						<form id="contactform" name="contactform"  @submit.prevent="processForm" method="POST">
							<div id="success"></div>
							<div class="row box-input">
								<div class="col-md-6 col-sm-6 col-xs-12 fix-left">
									<div class="form-group">
										<input type="text" name="clientname" v-model="clientname"  class="form-control" required="true" :placeholder="$lang.messages.yname" autocomplete="on">
									</div>
								</div>
								<div class="col-md-6 col-sm-6 col-xs-12 fix-right">
									<div class="form-group">
										<input type="email" name="email" v-model="email" :placeholder="$lang.messages.yemail" class="form-control">
									</div>
								</div>
							</div>
							<div class="form-group">
								<textarea class="form-control" name="msg" v-model="msg"  rows="2"  required="true" :placeholder="$lang.messages.ymsg" autocomplete="off"></textarea>
							</div>
							<button id="submit" type="submit" class="btn btn-info col-md-6 col-sm-6 col-xs-12">{{$lang.messages.send}}</button>
						</form>
					</div>
				</div>
			</div>
		</div>
		<b-modal id="myModalRef" ref="myModalRef" title="Feedback">
			<p class="my-4">{{resp}} {{err}}</p>
		</b-modal>
	</div>
</template>
<script>
	export default {
		name: 'contactform',
		data: function() {
			return {
				clientname: '',
				email: '',
				msg: '',
				err: '',
				resp: '',
				modalBody: null
				
			}
		},


		created() {
			if (this.resp!='') {
				this.modalBody = this.resp + '<br' + this.err;
			} else if (this.err!='') {
				this.modalBody = this.err;
			}
		},

	methods: {
		handleError: function(error) {
			this.err = this.strip(error);
		},

		handleResponse: function(response) {
			this.resp = response;
			console.log("resp");
			console.log(this.resp);
		},

		strip: function(html)
		{
			var tmp = document.createElement("DIV");
			tmp.innerHTML = html;
			return tmp.textContent || tmp.innerText || "";
		},

		processForm: function() {
			this.resp = '';
			this.err = '';
			console.log({ clientname: this.clientname, email: this.email, msg: this.msg, lang: this.$lang.getLang() });
			console.log(this);

			// let rawData = {
			// 	clientname: this.clientname,
			// 	email: this.email,
			// 	msg: this.msg,
			// 	lang: this.$lang.getLang()
			// };

			let rawData = this.msg;

			rawData = JSON.stringify(rawData);
			let formData = new FormData();
			formData.append('data', rawData);

			var self = this;	
			console.log(self);		
			this.$refs.myModalRef.show();
			this.$axios.post('localhost/api/contact', rawData, {headers: {'Content-Type': 'application/json'}})
			.then(response =>{
				this.handleResponse(response.data);
			})
			.catch(error =>{
				this.handleError(error);
			});
			this.msg = '';
			this.clientname = '';
			this.email = '';
		}
	}
}

</script>

<style scoped>
ul {
	list-style: initial;
}

ul li {
	display: block;
	line-height: 1.9em;
}

i.fas, i.fa, i.fab {
	padding: 0px 8px 0px 8px;
	max-width: 2em;
}

ul>li {
	padding-left: 0em;
}

@media (max-width: 768px) {  
	.jumbotron {
		padding: 0em 0em 2em 0em;
	}

	.jumbotron .display-4 {
		padding: 0.5em 0em 0.1em;
	}
}

</style>