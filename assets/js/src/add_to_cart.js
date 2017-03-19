$(function(){
  $('#add-to-cart-btn').click(function(e){
    e.preventDefault();
    var $sizes = $(this).parent().siblings('.sizes');
    var $sizeErrorMsg = $sizes.find('.error');
    $sizeErrorMsg.css('display', 'none');
    var $sizeOptions = $sizes.find('ul.active');
    var $selectedSize = $sizeOptions.find('li.selected div');
    if ($selectedSize.length == 0) return $sizeErrorMsg.show();

    var selectedSizeNumber = $selectedSize.text();
    var $otherColors = $(this).parent().siblings('.other-colors');

    var $selectedProduct = $otherColors.find('ul li.selected');
    var productID = $selectedProduct.data('product-id');

    axios.post('/cart/orders/' + productID, {
        size: selectedSizeNumber
      })
      .then(function (response) {
        ANIMATE.addAnimation($('#cart'), 'animated bounceIn');
        getCartProducts();
        // $('#add-to-cart-btn').html('<div class="fit-text">ADDED TO CART</div>');
      });
  });


  function getCartProducts(){
    axios.get('/cart/orders')
      .then(function (response) {
        $('#cart .products').remove();
        $('#cart').append(response.data);
      });
  }
});
