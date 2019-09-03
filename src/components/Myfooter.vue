<template>
    <footer class="py-4 bg-light">

        <!-- Copyright -->
        <div class="footer-copyright text-center py-3">2015 - {{ thisYear }}. <a href="http://blazhko.tech/">blazhko.tech</a>.
            None of the rights are reserved yet ¯\_(ツ)_/¯
        </div>
        <!-- Copyright -->

    </footer>
</template>
<script>
    export default {
        name: 'myfooter',
        computed: {
            thisYear: function () {
                let date = new Date();
                return date.getFullYear();
            }
        }
    }
</script>