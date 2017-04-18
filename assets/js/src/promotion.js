$(function(){
  var $promoCode = $('#promocode');
  var $promoSuccess = $promoCode.find('.promotion-success');
  var $promoError = $promoCode.find('.promotion-error');

  $('#product #promocode form').submit(function(e){
    e.preventDefault();
    var code = $(this).find('input').val();
    if (!code) return;
    resetMessages();
    axios.post('/promotions', {
        code: code
      })
      .then(function(response){
        $promoSuccess.show();
        ANIMATE.addAnimation($promoSuccess, 'animated bounceIn');
        $promoSuccess.find('span')
          .html(response.data.discount_percent + '%');
      })
      .catch(function(error){
        $promoError.show();
        ANIMATE.addAnimation($promoError, 'animated bounceIn');
      });
  });

  function resetMessages() {
    $promoSuccess.hide();
    $promoError.hide();
  }
});
