package article

import infra "arch/internal/infrastructure/article"

type Created struct {
	Article *infra.Entity
}

func (c *Created) Handle() {

}
