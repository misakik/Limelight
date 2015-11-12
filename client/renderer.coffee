search = $('input.search')
search.on 'keyup', ->
  console.log $(@).val()
