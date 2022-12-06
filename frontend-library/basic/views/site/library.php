<?php 
/** @var yii\web\View $this */

use yii\bootstrap5\LinkPager;

$this->title = 'Library';
$this->params['breadcrumbs'][] = $this->title;

include_once('../simple_html_dom.php');
?>

<h1>Library</h1>
<hr>

<ul>
    <div class="container">
        <div class='row'>
            <?php foreach($books as $book): ?>
                <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.2/font/bootstrap-icons.css">
                <div class='col-md-4' style='text-align: center;'>
                    <a href="<?= 'https://www.google.com/search?q=' . $book['title'] . ' ' . $book['author'] ?>">
                        <img src="<?= $book['thumbnailUrl'] ?>" class="rounded mx-auto" alt='<?=$book['title'] . ' Book Image'?>' title='<?=$book['title'] . ' Book Image'?>' width='150px' height='200px'>
                    </a>
                    <br>
                    <b><?= $book['title'] ?></b> <br>
                    <small><?= $book['author']?></small><br>
                    <small><?= ' Topics: ' . $book['categories'] ?></small><br>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <br><br>
                </div><br><br>
            <?php endforeach ?>
        </div>
    </div>
</ul>

<div style='
  margin: auto;
  width: 50%;
  padding: 10px;'><?= LinkPager::widget(['pagination' => $pagination]); ?> </div>