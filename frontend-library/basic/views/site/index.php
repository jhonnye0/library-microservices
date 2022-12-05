<?php

/** @var yii\web\View $this */

$this->title = 'E-Books Project';
?>
<div class="site-index">
    <br><br><br>
    <div class="jumbotron text-center bg-transparent">
        <h1 class="display-4">E-Books Software Reuse Project</h1>

        <p class="lead">Welcome to our Software Reuse Project E-Books!</p>

        <p><a class="btn btn-lg btn-success" href="index.php?r=site%2Flibrary">Access Books</a></p>
    </div>

    <div class="body-content">
        <br><br><br><br><br>
        <div class="row">
            <div class="col-lg-4">
                <center><h2>Books</h2></center>

                <p>Made with the intention of providing a library of digital
                     books for users to read or consult.</p>

                <p><a class="btn btn-outline-secondary" href="index.php?r=site%2Flibrary">Go to Library &raquo;</a></p>
            </div>
            <div class="col-lg-4">
                <center><h2>Technologies</h2></center>

                <p>Front-end based on Yii2 framework for PHP communicating
                     through API with Back-end based on components made in
                      Go language</p>

                <p><a class="btn btn-outline-secondary" href="https://github.com/jhonnye0/library-microservices">Github &raquo;</a></p>
            </div>
            <div class="col-lg-4">
                <h2>Github</h2>

                <div class="github-card" data-github="joaopedrobritot" data-width="400" data-height="" data-theme="default"></div>
                <script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

                <div class="github-card" data-github="jhonnye0" data-width="400" data-height="" data-theme="default"></div>
                <script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

                <!-- <p><a class="btn btn-outline-secondary" href="">Git &raquo;</a></p> -->
            </div>
        </div>

    </div>
</div>
