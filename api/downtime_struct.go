package api

type Recurrence struct {
	Period           *int     `json:"period,omitempty"`
	Type             *string  `json:"type,omitempty"`
	UntilDate        *int     `json:"until_date,omitempty"`
	UntilOccurrences *int     `json:"until_occurrences,omitempty"`
	WeekDays         []string `json:"week_days,omitempty"`
}

type DowntimeStruct []struct {
	Active      *bool       `json:"active,omitempty"`
	Canceled    *int        `json:"canceled,omitempty"`
	Disabled    *bool       `json:"disabled,omitempty"`
	End         *int        `json:"end,omitempty"`
	Id          *int        `json:"id,omitempty"`
	Message     *string     `json:"message,omitempty"`
	MonitorId   *int        `json:"monitor_id,omitempty"`
	MonitorTags []string    `json:"monitor_tags,omitempty"`
	ParentId    *int        `json:"parent_id,omitempty"`
	Timezone    *string     `json:"timezone,omitempty"`
	Recurrence  *Recurrence `json:"recurrence,omitempty"`
	Scope       []string    `json:"scope,omitempty"`
	Start       *int        `json:"start,omitempty"`
	CreatorID   *int        `json:"creator_id,omitempty"`
	UpdaterID   *int        `json:"updater_id,omitempty"`
	Type        *int        `json:"downtime_type,omitempty"`
}
