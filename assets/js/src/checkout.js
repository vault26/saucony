$(function(){
  registerShoppingCart();
  registerOrder();

  function registerShoppingCart() {
    showEmptyProducts();
    $('#shopping-cart .remove a').click(function(e){
      e.preventDefault();
      var productId = $(this).data('product-id');
      var size = $(this).data('product-size');
      var $productRow = $(this).parents('tr');
      alertify
        .okBtn('<div class="alert-btn fit-text danger">DELETE</div>')
        .cancelBtn('<div class="alert-btn fit-text">GO BACK</div>')
        .confirm('<div class="fit-text danger">CONFIRM TO DELETE THE ORDER?</div><div>ยืนยันที่จะลบรายการสินค้า</div>',
          function (e) {
            axios.delete('/cart/orders/' + productId + '?size=' + size)
              .then(function (response) {
                $productRow.remove();
                updateCheckoutOrders();
                showEmptyProducts();
              });
          },
          undefined
        );
    });

    $("#shopping-cart td.order-line-quantity a.remove").click(function(e){
      e.preventDefault();
      var quantity = parseInt($(this).siblings('.quantity').text());
      if (quantity == 1) {
        alertify
          .okBtn('<div class="alert-btn fit-text">OK</div>')
          .alert('<div class="fit-text">ITEM QUANTITY MUST BE GREATER THAN 1</div><div>ไม่สามารถลดจำนวนสินค้าได้น้อยกว่า 1</div>')
        return;
      }
      adjustOrder($(this), "remove", 1);
      $(this).blur();
    });

    $("#shopping-cart td.order-line-quantity a.add").click(function(e){
      e.preventDefault();
      adjustOrder($(this), "add", 1);
      $(this).blur();
    });
  }

  function showEmptyProducts() {
    if ($('#shopping-cart tbody tr.products').length == 0) {
      $('#shopping-cart tbody tr.no-products').show();
    }
  }

  function adjustOrder($link, operator, quantity) {
    var id = $link.data('product-id');
    var size = $link.data('product-size').toString();
    var $quantityComponent = $link.siblings('.quantity');
    var previousQuantity = parseInt($quantityComponent.text())

    axios.put('/cart/orders/' + id, {
      size: size,
      operator: operator,
      quantity: quantity
    })
    .then(function(response){
      updateCheckoutOrders();
    });
  }

  function updateCheckoutOrders() {
    axios.get('/cart/checkout_orders')
      .then(function(response){
        $('#checkout-orders').html(response.data);
        registerShoppingCart();
      });
  }

  function registerOrder() {
    $('#order-form button[type="submit"]').click(function(e){
      var $orderForm = $('#order-form');
      var form = {
        firstname: $orderForm.find('input#firstname').val(),
        lastname: $orderForm.find('input#lastname').val(),
        email: $orderForm.find('input#email').val(),
        phone: $orderForm.find('input#phone').val(),
        address: $orderForm.find('textarea#address').val()
      }
      var allInputsFilled = notEmpty(form.firstname) &&
        notEmpty(form.lastname) &&
        notEmpty(form.email) &&
        notEmpty(form.phone) &&
        notEmpty(form.address);
      if (!allInputsFilled) return;

      e.preventDefault();
      alertify
        .okBtn('<div class="alert-btn fit-text px-3 green">CONFIRM</div>')
        .cancelBtn('<div class="alert-btn secondary fit-text">CANCEL</div>')
        .confirm('<div class="fit-text green">CONFIRM ORDER</div><div>ยืนยันการสั่งซื้อสินค้า</div>',
          function (e) {
            $orderForm.submit();
          },
          undefined
        );
    })
  }

  function notEmpty(text) {
    return text.trim().length > 0
  }
});
