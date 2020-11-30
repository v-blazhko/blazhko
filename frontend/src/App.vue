<template>
    <div class="hello">
        <vue-ya-metrica :counter="48711056"
                        enable-hash
                        enable-webvisor></vue-ya-metrica>
        <navbar v-on:en="en" v-on:ru="ru"></navbar>
        <aboutme></aboutme>
        <div class="container my-5">
            <h1 class="text-center">{{$lang.messages.work}}</h1>
            <div class="container mt-1 mb-5">
                <div class="row">
                    <div class="col-md-8 offset-md-3 mx-auto">
                        <ul class="timeline">
                            <li>
                                <h6><a target="_blank" href="#">{{$lang.messages.exp.nokia.pos}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.exp.nokia.place}}</em></p>
                                <p><em href="#" class="">{{$lang.messages.exp.web.time}}</em></p>
                                <p class="pl-2">{{$lang.messages.exp.nokia.task1}}</p>
                                <p class="pl-2">{{$lang.messages.exp.nokia.task2}}</p>
                                <p class="pl-2">{{$lang.messages.exp.nokia.task3}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank" href="#">{{$lang.messages.exp.beauit.pos}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.exp.beauit.place}}</em></p>
                                <p><em href="#" class="">{{$lang.messages.exp.web.time}}</em></p>
                                <p class="pl-2">{{$lang.messages.exp.beauit.task1}}</p>
                                <p class="pl-2">{{$lang.messages.exp.beauit.task2}}</p>
                                <p class="pl-2">{{$lang.messages.exp.beauit.task3}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank" href="#">{{$lang.messages.exp.web.pos}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.exp.web.place}}</em></p>
                                <p><em href="#" class="">{{$lang.messages.exp.web.time}}</em></p>
                                <p class="pl-2">{{$lang.messages.exp.web.task1}}</p>
                                <p class="pl-2">{{$lang.messages.exp.web.task2}}</p>
                                <p class="pl-2">{{$lang.messages.exp.web.task3}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank" href="#">{{$lang.messages.exp.emb.pos}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.exp.emb.place}}</em></p>
                                <p><em href="#" class="">{{$lang.messages.exp.emb.time}}</em></p>
                                <p class="pl-2">{{$lang.messages.exp.emb.task1}}</p>
                                <p class="pl-2">{{$lang.messages.exp.emb.task2}}</p>
                                <p class="pl-2">{{$lang.messages.exp.emb.task3}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank" href="#">{{$lang.messages.exp.fwd.pos}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.exp.fwd.place}}</em></p>
                                <p><em href="#" class="">{{$lang.messages.exp.fwd.time}}</em></p>
                                <p class="pl-2">{{$lang.messages.exp.fwd.task1}}</p>
                                <p class="pl-2">{{$lang.messages.exp.fwd.task2}}</p>
                                <p class="pl-2">{{$lang.messages.exp.fwd.task3}}</p>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>

            <h1 class="text-center">{{$lang.messages.skill}}</h1>
            <div class="container-fluid skills">
                <skill cskill="Golang" cclass="greenish"></skill>
                <skill cskill="Python 3" cclass=""></skill>
                <skill cskill="Javascript" cclass="greenish"></skill>
                <skill cskill="ES6" cclass="greenish"></skill>
                <skill cskill="Databases" cclass="greyish"></skill>
                <skill cskill="Docker" cclass=""></skill>
                <skill cskill="ES6" cclass=""></skill>
                <skill cskill="Vue.js" cclass="greenish"></skill>
                <skill cskill="Adobe Illustrator" cclass=""></skill>
                <skill cskill="Matlab" cclass=""></skill>
                <skill cskill="Docker" cclass=""></skill>
                <skill cskill="Cisco" cclass=""></skill>
                <skill cskill="Embedded systems" cclass=""></skill>
                <skill cskill="Network administration" cclass="greyish"></skill>
                <skill cskill="C/C++" cclass="greenish"></skill>
                <skill cskill="PHP" cclass=""></skill>
                <skill cskill="Linux" cclass="greyish"></skill>
                <skill cskill="AutoIT" cclass="greenish"></skill>
                <skill cskill="Node.js" cclass="greenish"></skill>
                <skill cskill="CI/CD" cclass="greenish"></skill>
                <skill cskill="UX/UI" cclass="greenish"></skill>
            </div>

            <h1 class="text-center">{{$lang.messages.education}}</h1>
            <div class="container mt-1 mb-5">
                <div class="">
                    <div class="col-md-7 offset-md-3 mx-auto">
                        <ul class="timeline">
                            <li>
                                <h6><a target="_blank"
                                       href="http://tuni.fi">{{$lang.messages.edu.master.name}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.edu.master.time}}</em></p>
                                <p>{{$lang.messages.edu.master.descr}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank"
                                       href="http://www.tut.fi/en/home">{{$lang.messages.edu.tut.name}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.edu.tut.time}}</em></p>
                                <p>{{$lang.messages.edu.tut.descr}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank" :href="$lang.messages.edu.sut.link">{{$lang.messages.edu.sut.name}}</a>
                                </h6>
                                <p><em href="#" class="">{{$lang.messages.edu.sut.time}}</em></p>
                                <p>{{$lang.messages.edu.sut.descr}}</p>
                            </li>
                            <li>
                                <h6><a target="_blank"
                                       href="https://www.netacad.com/">{{$lang.messages.edu.cisco.name}}</a></h6>
                                <p><em href="#" class="">{{$lang.messages.edu.cisco.time}}</em></p>
                                <p>{{$lang.messages.edu.cisco.descr}}</p>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <contactform></contactform>
        </div>
        <myfooter></myfooter>
    </div>
</template>

<script>
    export default {
        name: 'app',
        data() {
            return {
                msg: 'Welcome to Your Vue.js App',
                counter: 0,
                publicPath: process.env.BASE_URL
            }
        },

        metaInfo: {
            meta: [
                {charset: 'utf-8'},
                {name: 'viewport', content: 'width=device-width, initial-scale=1'}
            ]
        },

        methods: {
            ru() {
                this.$lang.setLang('ru')
            },
            en() {
                this.$lang.setLang('en')
            }
        },
    }
</script>

<style>

    @import "https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css";
    @import "https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.0/animate.css";
    @import "https://kit.fontawesome.com/e80a83c8f1.js";

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
        border: 3px solid #17a2b8;
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
