package weapon

import (
	"time"

	"github.com/yohamta/godanmaku/danmaku/internal/shot"
)

// Normal represents player's weapon
type Normal struct {
	lastShotTime time.Time
	shotKind     shot.Kind
}

// NewNormal creates new struct
func NewNormal(shotKind shot.Kind) *Normal {
	w := &Normal{}
	w.shotKind = shotKind

	return w
}

// Fire create shots
func (w *Normal) Fire(shooter Shooter, shots []*shot.Shot) {
	if time.Since(w.lastShotTime).Milliseconds() < 350 {
		return
	}
	w.lastShotTime = time.Now()

	for i := 0; i < len(shots); i++ {
		s := shots[i]
		if s.IsActive() {
			continue
		}
		s.Init(w.shotKind, shooter.GetDegree())
		s.SetPosition(shooter.GetX(), shooter.GetY())
		break
	}

}