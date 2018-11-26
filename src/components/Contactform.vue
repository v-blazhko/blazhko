<template>
	<div id="contact">
		<h1 class="text-center">{{$lang.messages.sendmsg}}</h1>
		<div class="row">
			<div class="col-sm-12 col-lg-4">
				<ul class="row" style="display:flex;">
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fas fa-map-marker-alt mr-2"></i>   Location: Tampere</li>
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fa fa-envelope mr-2"></i>   Email: veronica@blazhko.tech</li>
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-vk mr-2"></i>   VK: vk.com/v_blazhko</li>
					<li  class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-linkedin-in mr-2"></i>   LinkedIn: my account</li>			
					<li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-github mr-2"></i>   Github: my account</li>
				</ul>
			</div>
			<div class="col-sm-12 col-lg-8">
				<div class="row">
					<div class="col-md-12 mx-auto">
						<form id="contactform" name="contactform" method="POST">
							<div id="success"></div>
							<div class="row box-input">
								<div class="col-md-6 col-sm-6 col-xs-12 fix-left">
									<div class="form-group">
										<input type="text" name="clientname" class="form-control" id="name" required="true" placeholder="Your name" autocomplete="off">
									</div>
								</div>
								<div class="col-md-6 col-sm-6 col-xs-12 fix-right">
									<div class="form-group">
										<input type="email" name="email" class="form-control" id="email"  required="true" placeholder="E-mail" autocomplete="off">
									</div>
								</div>
							</div>
							<div class="form-group">
								<textarea class="form-control" name="msg" rows="2"  required="true" placeholder="Your message" autocomplete="off"></textarea>
							</div>
							<button id="submit" type="submit" class="btn btn-info">{{$lang.messages.send}}</button>
						</form>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	name: 'contactform'
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
</style>