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
                <?php
                    $search_keyword = str_replace(' ','+',$book->title . ' ' . $book->author . ' book');
                    $newhtml = file_get_html("https://www.google.com/search?q=".$search_keyword."&tbm=isch");
                    // $newhtml = file_get_html("https://www.google.com/search?q=" . $search_keyword . "&tbm=bks&source=lnt&tbs=bkt:b&sa=X&ved=2ahUKEwjKv7vVneP7AhV4qZUCHYbTANkQpwV6BAgBECA&biw=1366&bih=671&dpr=1");
                    for($i = 0; $i <= count($newhtml->find('img')); $i++)
                    {
                        if(!is_null($newhtml->find('img', $i)))
                        {
                            $choosenImg = $newhtml->find('img', $i);
                            if(str_contains($choosenImg->src, 'data:image/jpeg'))// and str_contains($choosenImg->alt, 'Amazon'))// and str_contains($choosenImg->id, 'dimg_') and str_contains($choosenImg->height, '156'))
                                break;
                        }
                        
                    }
                    
                    $result_image_source = $choosenImg->src;
                ?>
                <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.2/font/bootstrap-icons.css">
                <div class='col-md-4' style='text-align: center;'>
                    <a href="<?= 'https://www.google.com/search?q=' . $search_keyword ?>">
                        <img src="<?= $result_image_source ?>" class="rounded mx-auto" alt='<?=$book->title . ' Book Image'?>' title='<?=$book->title . ' Book Image'?>' width='150px' height='200px'>
                    </a>
                    <br>
                    <b><?= $book->title ?></b> <br>
                    <small><?= $book->author?></small><br>
                    <small><?= ' Genres: ' . $book->topic ?></small><br>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                    <i class='bi bi-star-fill'></i>
                </div><br>
            <?php endforeach ?>
        </div>
    </div>
</ul>

<?= LinkPager::widget(['pagination' => $pagination
                        // 'pageCssClass' => 'btn btn-warning',
                        // 'nextPageCssClass' => 'btn btn-dark',
                        // 'prevPageCssClass' => 'btn btn-dark',
                        // 'activePageCssClass' => 'btn btn-danger link-light',
                        // 'options' => [
                        //     'class' => 'ip-mosaic__pagin-list',
                        // ]
                        ]); ?>