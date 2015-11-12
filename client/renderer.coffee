search = $('input.search')
search.on 'keyup', ->
  $.ajax
    url: 'http://localhost:8000/search/' + $(@).val()
    dataType: 'json'
    success: (data) ->
      console.log data
