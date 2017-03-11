$(function(){
  var $product = $('#product');
  var $mainImage = $product.find('.main-image .bg-img');
  var $sideImage = $product.find('.side-images');
  var $sideImageOptions = $sideImage.find('li');

  var $otherColor = $('#product .other-colors');
  var $otherColorOptions = $otherColor.find('li');
  var $colorSelectionText = $otherColor.find('.selected-color');

  var $size = $('#product .sizes');
  var $sizeOptions = $size.find('li');
  var $sizeSelectionText = $size.find('.selected-size');

  $sideImageOptions.hover(sideImagesHoverHandler);
  $otherColorOptions.click(otherColorClickHandler);
  $sizeOptions.click(sizeClickHandler);

  function clearOtherSelection(options, className){
    options.each(function(index, item){
      $(item).removeClass(className);
    });
  }

  function sideImagesHoverHandler(){
    clearOtherSelection($(this).siblings(), 'selected');
    $(this).addClass('selected');
    var newImgName = $(this).find('.bg-img').data('img');
    setMainImage(newImgName);
  }

  function setMainImage(imgName) {
    var mainImgUrl = $mainImage.css('background-image');
    var mainImgName = $mainImage.data('img');
    var imgRegExp = new RegExp(mainImgName, 'g');
    var newImgUrl = mainImgUrl.replace(imgRegExp, imgName);
    $mainImage.css('background-image', newImgUrl);
    $mainImage.data('img', imgName);
  }

  function otherColorClickHandler(){
    $(this).addClass('selected');
    clearOtherSelection($(this).siblings(), 'selected');
    $colorSelectionText.text($(this).data('color'));

    var newImgName = $(this).find('.bg-img').data('img');
    setMainImage(newImgName);
    var productID = $(this).data('product-id');
    toggleSideImages(productID);
    toggleSize(productID);
  }

  function toggleSideImages(productId){
    var $productSideImages = $('#side-images-' + productId);
    $productSideImages.addClass('active');
    clearOtherSelection($productSideImages.siblings(), 'active');
    loadSideImages($productSideImages)
  }

  function loadSideImages($sideImage) {
    $sideImage.find('.bg-img').each(function(index, item){
      // background style not set
      if (!$(this).attr('style')) {
        $(this).attr('style', $(this).data('img-url'));
      }
    })
  }

  function toggleSize(productId) {
    var $productSideImages = $('#size-' + productId);
    $productSideImages.addClass('active');
    clearOtherSelection($productSideImages.siblings(), 'active');
    setSizeSelectionText();
  }

  function sizeClickHandler(){
    $(this).addClass("selected");
    clearOtherSelection($(this).siblings(), 'selected');
    setSizeSelectionText();
  }

  function setSizeSelectionText() {
    var $selectedSizeOption = $size.find('ul.active li.selected');
    if ($selectedSizeOption.length == 0) {
      $sizeSelectionText.text('-');
      return
    }
    $sizeSelectionText.text($selectedSizeOption.data('size'));
  }
});
