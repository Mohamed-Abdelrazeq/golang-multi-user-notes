package supa

import (
	"context"

	"github.com/nedpals/supabase-go"
)

var SupabaseClient *supabase.Client
var ctx context.Context

func OpenClientConnection() {
	supabaseUrl := "https://zamgxydgfozcainpiryh.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InphbWd4eWRnZm96Y2FpbnBpcnloIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTMyODIxNTksImV4cCI6MjAwODg1ODE1OX0.9bVzAhVk9RPGqy2wr6XH_vgwnQMrKLuxjfHUOMtX4d0"
	SupabaseClient = supabase.CreateClient(supabaseUrl, supabaseKey)
	ctx = context.Background()
}
