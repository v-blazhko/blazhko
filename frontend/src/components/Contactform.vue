<template>
  <div id="contact">
    <h1 class="text-center">{{ $lang.messages.sendmsg }}</h1>
    <div class="row">
      <div class="col-sm-12 col-lg-4">
        <ul class="row" style="display:flex;">
          <li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fas fa-map-marker-alt mr-2"></i>Location: <a
              href="https://www.google.fi/maps/place/%D0%A2%D0%B0%D0%BC%D0%BF%D0%B5%D1%80%D0%B5/@61.6319675,23.5501175,10z/data=!4m5!3m4!1s0x468edf554593da5d:0x6adfe3bd1e0b22c0!8m2!3d61.4977524!4d23.7609535"
              target="_blank">Tampere</a></li>
          <li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fa fa-envelope mr-2"></i>Email: <a
              href="mailto:veronica@blazhko.tech" target="_blank">veronica@blazhko.tech</a></li>
          <li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-linkedin-in mr-2"></i>LinkedIn: <a
              href="https://linkedin.com/in/v-blazhko/" target="_blank">v-blazhko</a></li>
          <li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-github mr-2"></i>Github: <a
              href="https://github.com/v-blazhko/" target="_blank">v-blazhko</a></li>
          <li class="col-lg-12 col-md-6 col-sm-6 col-xs-6"><i class="fab fa-vk mr-2"></i>VK: <a
              href="https://vk.com/v_blazhko" target="_blank">v_blazhko</a></li>
        </ul>
      </div>
      <div class="col-sm-12 col-lg-8">
        <div class="row">
          <div class="col-md-12 mx-auto">
            <form id="contactform" name="contactform" @submit.prevent="processForm" method="POST">
              <div id="success"></div>
              <div class="row box-input">
                <div class="col-md-6 col-sm-6 col-xs-12 fix-left">
                  <div class="form-group">
                    <input type="text" name="name" v-model="name" class="form-control" required="true"
                           :placeholder="$lang.messages.yname" autocomplete="on">
                  </div>
                </div>
                <div class="col-md-6 col-sm-6 col-xs-12 fix-right">
                  <div class="form-group">
                    <input type="email" name="email" v-model="email" :placeholder="$lang.messages.yemail"
                           class="form-control">
                  </div>
                </div>
              </div>
              <div class="form-group">
                <textarea class="form-control" name="message" v-model="message" rows="2" required="true"
                          :placeholder="$lang.messages.ymsg" autocomplete="off"></textarea>
              </div>
              <button id="submit" type="submit" class="btn btn-info col-md-6 col-sm-6 col-xs-12">
                {{ $lang.messages.send }}
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <b-modal id="myModalRef" ref="myModalRef" title="Feedback">
      <p class="my-4">{{ resp }} {{ err }}</p>
    </b-modal>
  </div>
</template>
<script>
export default {
  name: 'contactform',
  data: function () {
    return {
      name: '',
      email: '',
      message: '',
      err: '',
      resp: '',
      modalBody: null
    }
  },

  created() {
    if (this.resp !== '') {
      this.modalBody = this.resp + '<br' + this.err;
    } else if (this.err !== '') {
      this.modalBody = this.err;
    }
  },

  methods: {
    handleError: function () {
      this.err = "An error occurred while submitting your response";
    },

    handleResponse: function (responseCode) {
      if ((responseCode < 300) && (responseCode > 199)) {
        this.resp = "Thanks for submitting your message! I will reach out to you as fast as possible :)"
      } else {
        this.resp = "An error occurred while submitting your response"
      }
      this.$refs.myModalRef.show();
    },

    strip: function (html) {
      let tmp = document.createElement("div");
      tmp.innerHTML = html;
      return tmp.textContent || tmp.innerText || "";
    },

    processForm: function () {
      this.resp = '';
      this.err = '';
      let rawData = {
        name: this.name,
        email: this.email,
        message: this.message,
        lang: this.$lang.getLang()
      };

      rawData = JSON.stringify(rawData);
      let formData = new FormData();
      formData.append('data', rawData);

      this.$axios.post('/api/contact', rawData, {headers: {'Content-Type': 'application/json'}})
          .then(response => {
            this.handleResponse(response.status);
          })
          .catch(error => {
            this.handleError(error);
          });
      this.message = '';
      this.name = '';
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

ul > li {
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