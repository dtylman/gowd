package bootstrap

import (
	"github.com/dtylman/gowd"
)

/*
<nav class="navbar navbar-default">
  <div class="container-fluid">
      <ul class="nav navbar-nav">
        <button type="button" class="btn btn-default navbar-btn">Sign in</button>
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="#">Action</a></li>
            <li><a href="#">Another action</a></li>
            <li><a href="#">Something else here</a></li>
            <li role="separator" class="divider"></li>
            <li><a href="#">Separated link</a></li>
          </ul>
        </li>
      </ul>

  </div><!-- /.container-fluid -->
</nav>
*/

const NavbarDefault = "navbar-default"

type Navbar struct {
	*gowd.Element
	Lists     []*List
	container *gowd.Element
}

func NewNavBar(navbarType string) *Navbar {
	nb := &Navbar{}
	nb.Element = NewElement("nav", navbarType)
	nb.container = NewContainer(true)
	nb.Element.AddElement(nb.container)
	nb.Lists = make([]*List, 0)
	return nb
}

func (nb*Navbar) AddList() *List {
	list := NewList(ListUnordered, "nav navbar-nav nav-pills")
	nb.Lists = append(nb.Lists, list)
	nb.container.AddElement(list.Element)

	return list
}