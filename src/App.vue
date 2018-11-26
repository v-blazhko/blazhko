<template>
  <div class="hello">
    <navbar v-on:en="en" v-on:ru="ru"></navbar> 
    <aboutme></aboutme>
    <div class="container my-5">

      <h1 class="text-center">{{$lang.messages.work}}</h1>

      <div class="container mt-5 mb-5">
        <div class="row">
          <div class="col-md-8 offset-md-3 mx-auto">
            <h4>Latest News</h4>
            <ul class="timeline">
              <li>
                <a target="_blank" href="https://www.totoprayogo.com/#">New Web Design</a>
                <a href="#" class="float-right">21 March, 2014</a>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque scelerisque diam non nisi semper, et elementum lorem ornare. Maecenas placerat facilisis mollis. Duis sagittis ligula in sodales vehicula....</p>
              </li>
              <li>
                <a href="#">21 000 Job Seekers</a>
                <a href="#" class="float-right">4 March, 2014</a>
                <p>Curabitur purus sem, malesuada eu luctus eget, suscipit sed turpis. Nam pellentesque felis vitae justo accumsan, sed semper nisi sollicitudin...</p>
              </li>
              <li>
                <a href="#">Awesome Employers</a>
                <a href="#" class="float-right">1 April, 2014</a>
                <p>Fusce ullamcorper ligula sit amet quam accumsan aliquet. Sed nulla odio, tincidunt vitae nunc vitae, mollis pharetra velit. Sed nec tempor nibh...</p>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <h1 class="text-center">{{$lang.messages.skill}}</h1>
      <div class="container-fluid skills">
        <skill cskill="Javascript" cclass="greenish"></skill>
        <skill cskill="ES6" cclass="greenish"></skill>
        <skill cskill="PHP7" cclass="greenish"></skill>
        <skill cskill="MySQL" cclass="greyish"></skill>       
        <skill cskill="ES 6" cclass=""></skill>
        <skill cskill="Vue.js" cclass="greenish"></skill>
        <skill cskill="CSS3" cclass=""></skill>
        <skill cskill="LESS" cclass=""></skill>
        <skill cskill="HTML5" cclass=""></skill>
        <skill cskill="Adobe Illustrator" cclass=""></skill>
        <skill cskill="Python3" cclass="greenish"></skill>
        <skill cskill="Matlab" cclass=""></skill>
        <skill cskill="Windows" cclass="greyish"></skill>
        <skill cskill="CMS Wordpres" cclass=""></skill>

        <skill cskill="Cisco" cclass=""></skill>
        <skill cskill="Embedded systems" cclass=""></skill>

        <skill cskill="Network administration" cclass="greyish"></skill>
        <skill cskill="C/C++" cclass="greenish"></skill>

        <skill cskill="Adobe Photoshop" cclass=""></skill>

        <skill cskill="Microcontrollers" cclass="greenish"></skill>

        <skill cskill="Raspberry PI" cclass=""></skill>

        <skill cskill="Arduino" cclass=""></skill>

        <skill cskill="Orange PI" cclass=""></skill>

        <skill cskill="Linux" cclass="greyish"></skill>
        <skill cskill="AutoIT" cclass="greenish"></skill>
        <skill cskill="Bootstrap 4" cclass=""></skill>
        <skill cskill="CMS Joomla" cclass=""></skill>
        <skill cskill="Node.js" cclass="greenish"></skill>

        <skill cskill="webpack" cclass="greenish"></skill>

        <skill cskill="Apache" cclass="greyish"></skill>

        <skill cskill="nginx" cclass="greyish"></skill>
         <skill cskill="UX/UI Design" cclass="greenish"></skill>

      </div>
      <h1 class="text-center">{{$lang.messages.education}}</h1>


      <div class="container mt-5 mb-5">
        <div class="row">
          <div class="col-md-8 offset-md-3 mx-auto">
            <h4>Latest News</h4>
            <ul class="timeline">
              <li>
                <a target="_blank" href="https://www.totoprayogo.com/#">New Web Design</a>
                <a href="#" class="float-right">21 March, 2014</a>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque scelerisque diam non nisi semper, et elementum lorem ornare. Maecenas placerat facilisis mollis. Duis sagittis ligula in sodales vehicula....</p>
              </li>
              <li>
                <a href="#">21 000 Job Seekers</a>
                <a href="#" class="float-right">4 March, 2014</a>
                <p>Curabitur purus sem, malesuada eu luctus eget, suscipit sed turpis. Nam pellentesque felis vitae justo accumsan, sed semper nisi sollicitudin...</p>
              </li>
              <li>
                <a href="#">Awesome Employers</a>
                <a href="#" class="float-right">1 April, 2014</a>
                <p>Fusce ullamcorper ligula sit amet quam accumsan aliquet. Sed nulla odio, tincidunt vitae nunc vitae, mollis pharetra velit. Sed nec tempor nibh...</p>
              </li>
            </ul>
          </div>
        </div>
      </div>


      <contactform></contactform>

      <game></game>






    </div>

    <myfooter></myfooter>
  </div>
</template>

<script>

export default {
  name: 'app',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      counter: 0
    }
  },

  methods: {
    ru(){
      this.$lang.setLang('ru')
    },
    en(){
      this.$lang.setLang('en')
    }  
  }
}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style >

@import "https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css";
@import "https://use.fontawesome.com/releases/v5.5.0/css/all.css";
h1, h2 {
  font-weight: normal;
  padding: 1em 0em;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

ul.timeline {
  list-style-type: none;
  position: relative;
}
ul.timeline:before {
  content: ' ';
  background: #d4d9df;
  display: inline-block;
  position: absolute;
  left: 0px;
  width: 2px;
  height: 100%;
  z-index: 400;
}
ul.timeline > li {
  margin: 20px 0;
  padding-left: 20px;
}
ul.timeline > li:before {
  content: ' ';
  background: white;
  display: inline-block;
  position: absolute;
  border-radius: 50%;
  border: 3px solid #22c0e8;
  left: -10px;
  width: 20px;
  height: 20px;
  z-index: 400;
}

.skills {
  display: flex;
  flex-wrap: wrap;
}


</style>
